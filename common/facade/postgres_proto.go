package facade

import (
	"errors"
	"fmt"
	"strings"

	"github.com/forhsd/logger"
	"github.com/ouseikou/sqlbuilder/common/clause"
	pb "github.com/ouseikou/sqlbuilder/gen/proto"
	"github.com/ouseikou/sqlbuilder/util"
	xorm "xorm.io/builder"
)

// PostgresModelBuilderFacade 具体Facade实现
type PostgresModelBuilderFacade struct {
	AbstractModelBuilderFacade
}

func (f *PostgresModelBuilderFacade) BuildSQL(request *pb.BuilderRequest) (*xorm.Builder, error) {
	builders := request.Builders
	// 目前只有一层
	if len(builders) != 1 {
		return nil, errors.New("SQL Builder 暂时只支持一层")
	}
	// 上游确定是 model 构造器
	sqlRef := builders[0].Sql.GetModel()
	// 选择方言
	dialectBuilder, err := f.BuildDialect(request)
	if err != nil {
		return nil, err
	}
	// builder(selects, from , as)
	basicBuilder := f.BuildBasic(dialectBuilder, sqlRef)
	// 调用BuildOther方法处理其他逻辑
	builder := f.BuildOther(basicBuilder, sqlRef)

	return builder, nil
}

func (f *PostgresModelBuilderFacade) BuildDialect(request *pb.BuilderRequest) (*xorm.Builder, error) {
	logger.Debug("[模板方法模式]-[postgres] 实现: 构建方言 builder...")
	return f.AbstractModelBuilderFacade.BuildDialect(request)
}

func (f *PostgresModelBuilderFacade) BuildBasic(dialectBuilder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	logger.Debug("[模板方法模式]-[postgres] 实现: 构建 builder 基础内容...")

	selects, tableSchema, fromAlias := convert2BasicArgsByProto(sqlRef)

	basicBuilder := buildBasicByProto(dialectBuilder, selects, tableSchema, fromAlias)

	return basicBuilder
}

func buildBasicByProto(
	dialectBuilder *xorm.Builder,
	selects []string,
	sub string,
	fromAlias string) *xorm.Builder {
	// XORM 只有Select不能构造SQL, 至少需要 Select && From
	basicBuilder := dialectBuilder.Select(selects...).From(sub, fromAlias)
	return basicBuilder
}

func convert2BasicArgsByProto(sqlRef *pb.SQLReference) ([]string, string, string) {
	selectClause := sqlRef.Select

	// 类型断言: Column 或者 Expression
	selects := extraColumnOrExpressionStringsByProto(selectClause)

	fromClause := sqlRef.From

	tableSchema := fmt.Sprintf(
		clause.PGTableColumn,
		fromClause.TableSchema,
		fromClause.TableName)

	return selects, tableSchema, fmt.Sprintf(clause.PGStringLiteralSafe, fromClause.TableAlias)
}

func extraColumnOrExpressionStringsByProto(mixClause []*pb.MixField) []string {
	var selects []string
	for _, mixItem := range mixClause {
		// 使用类型断言判断具体是 Column 还是 Expression
		switch v := mixItem.GetMix().(type) {
		case *pb.MixField_Column:
			column := v.Column
			formatResult := util.Ternary(column.UseAs, clause.PGTableColumnAs, clause.PGTableColumn)
			format := formatResult.(string)
			selects = append(
				selects,
				calAppendTextByColumnAsByProto(format, column, column.UseAs, column.Alias))
		case *pb.MixField_Expression:
			expression := v.Expression
			vars2Strings := calExpressionVarsFormatStringsByProto(expression)
			format := getExpressionFuncFormatByProto(expression)
			selects = append(
				selects,
				fmt.Sprintf(
					format,
					vars2Strings...))
		default:
			// 处理未知类型
			fmt.Printf("Unexpected type: %T\n", v)
		}
	}
	return selects
}

func calAppendTextByColumnAsByProto(format string, column *pb.Column, userAs bool, alias string) string {
	if userAs {
		return fmt.Sprintf(format, column.Table, column.Field, alias)
	}
	return fmt.Sprintf(format, column.Table, column.Field)
}

func getExpressionFuncFormatByProto(expression *pb.Expression) string {
	callType := expression.CallType
	if expression.UseAs {
		switch callType {
		case pb.CallType_CALL_TYPE_AGG:
			return clause.PGAggregationAsFuncFormatMap[clause.AggFunc(expression.Call)]
		case pb.CallType_CALL_TYPE_INNER:
			return clause.InnerAsFuncFormatMap[clause.InnerFunc(expression.Call)]
		default:
			return ""
		}
	}

	switch callType {
	case pb.CallType_CALL_TYPE_AGG:
		return clause.AggregationFuncFormatMap[clause.AggFunc(expression.Call)]
	case pb.CallType_CALL_TYPE_INNER:
		return clause.InnerFuncFormatMap[clause.InnerFunc(expression.Call)]
	default:
		return ""
	}
}

func calExpressionVarsFormatStringsByProto(expr *pb.Expression) []interface{} {
	vars := expr.Vars
	// 表达式参数: int, string, clause.Column
	var expressions []interface{}
	for _, varItem := range vars {
		switch v := varItem.GetVars().(type) {
		case *pb.MixVars_Column:
			column := v.Column
			formatResult := util.Ternary(column.UseAs, clause.PGTableColumnAs, clause.PGTableColumn)
			format := formatResult.(string)
			expressions = append(
				expressions,
				calAppendTextByColumnAsByProto(format, column, column.UseAs, expr.CallAs))
		case *pb.MixVars_Context:
			exprString := v.Context
			expressions = append(expressions, exprString)
		case *pb.MixVars_Number:
			exprInt := v.Number
			expressions = append(
				expressions,
				fmt.Sprintf(
					clause.AnyLiteral,
					exprInt))
		default:
			continue
		}
	}

	if expr.UseAs {
		expressions = append(expressions, expr.CallAs)
	}

	return expressions
}

func (f *PostgresModelBuilderFacade) BuildOther(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	logger.Debug("[模板方法模式]-[postgres] 实现: 构建 builder 其他内容...")

	// 判断是否组装 Join
	if len(sqlRef.Join) > 0 {
		builder = buildJoinByProto(builder, sqlRef)
	}

	// 判断是否组装 Where
	if len(sqlRef.Where) > 0 {
		builder = buildWhereByProto(builder, sqlRef)
	}

	// 判断是否组装 GroupBy
	if len(sqlRef.GroupBy) > 0 {
		builder = buildGroupByByProto(builder, sqlRef)
	}

	// 判断是否组装 OrderBy
	if len(sqlRef.OrderBy) > 0 {
		builder = buildOrderByProto(builder, sqlRef)
	}

	// 判断是否组装 Limit
	if &sqlRef.Limit != nil {
		builder = buildLimitProto(builder, sqlRef)
	}

	return builder
}

func buildJoinByProto(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	return builder
}
func buildWhereByProto(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	return builder
}
func buildOrderByProto(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	return builder
}
func buildLimitProto(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	return builder
}

func buildGroupByByProto(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	groupByClause := sqlRef.GroupBy
	// 类型断言: Column 或者 Expression
	groupByItems := extraColumnOrExpressionStringsByProto(groupByClause)

	// 使用 strings.Join 来合并切片元素，以逗号分隔
	joinedItems := strings.Join(groupByItems, clause.COMMA)

	// xorm 的 GroupBy 方法只接收一个字符串参数，因此要自己将多个 GroupBy 拼接为1个
	groupByFragment := fmt.Sprintf(clause.StringLiteral, joinedItems)

	return builder.GroupBy(groupByFragment)
}
