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

func (f *PostgresModelBuilderFacade) BuildBasic(dialectBuilder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	logger.Debug("[模板方法模式]-[postgres] 实现: 构建 builder 基础内容...")

	selects, tableSchema, fromAlias := createBasicArgsByProto(sqlRef)

	basicBuilder := buildBasicByProto(dialectBuilder, selects, tableSchema, fromAlias)

	return basicBuilder
}

// buildBasicByProto 构建 执行xorm sqlbuilder必须的select和from
func buildBasicByProto(
	dialectBuilder *xorm.Builder,
	selects []string,
	sub string,
	fromAlias string) *xorm.Builder {
	// XORM 只有Select不能构造SQL, 至少需要 Select && From
	basicBuilder := dialectBuilder.Select(selects...).From(sub, fromAlias)
	return basicBuilder
}

// createBasicArgsByProto 提取构建 buildBasicByProto 必须的参数
func createBasicArgsByProto(sqlRef *pb.SqlReference) ([]string, string, string) {
	selectClause := sqlRef.Select

	// 类型断言: Column 或者 Expression
	selects := extraColumnOrExpressionStringsByProto(selectClause)

	fromClause := sqlRef.From

	tableSchema := fmt.Sprintf(clause.PGTableColumn, fromClause.TableSchema, fromClause.TableName)

	return selects, tableSchema, fmt.Sprintf(clause.PGStringLiteralSafe, fromClause.TableAlias)
}

// extraColumnOrExpressionStringsByProto 将MixField转换为对应的string字符串
func extraColumnOrExpressionStringsByProto(mixClause []*pb.MixField) []string {
	var selects []string
	for _, mixItem := range mixClause {
		//switch v := mixItem.GetMix().(type) {
		//case *pb.MixField_Column:
		//	column := v.Column
		//	formatResult := util.Ternary(column.UseAs, clause.PGTableColumnAs, clause.PGTableColumn)
		//	format := formatResult.(string)
		//	selects = append(
		//		selects,
		//		calAppendTextByColumnAsByProto(format, column, column.UseAs, column.Alias))
		//case *pb.MixField_Expression:
		//	expression := v.Expression
		//	vars2Strings := calExpressionVarsFormatStringsByProto(expression)
		//	format := getExpressionFuncFormatByProto(expression)
		//	selects = append(
		//		selects,
		//		fmt.Sprintf(
		//			format,
		//			vars2Strings...))
		//default:
		//	// 处理未知类型
		//	fmt.Printf("Unexpected type: %T\n", v)
		//}

		result := processMixFieldItemByProto(mixItem)
		if result != "" {
			selects = append(selects, result)
		}
	}
	return selects
}

func processMixFieldItemByProto(mixItem *pb.MixField) string {
	// 使用类型断言判断具体是 Column 还是 Expression
	switch v := mixItem.GetMix().(type) {
	case *pb.MixField_Column:
		column := v.Column
		format := util.Ternary(column.UseAs, clause.PGTableColumnAs, clause.PGTableColumn).(string)
		return formatColumnByProto(column, format)
	case *pb.MixField_Expression:
		expression := v.Expression
		return formatExpressionByProto(expression)
	default:
		return ""
	}
}

// formatColumnByProto 构建 Column 类型的格式化字符串填充
func formatColumnByProto(column *pb.Column, format string) string {
	return calAppendTextByColumnAsByProto(format, column, column.UseAs, column.Alias)
}

// formatExpressionByProto 构建 Expression 类型的格式化字符串填充
func formatExpressionByProto(expression *pb.Expression) string {
	vars2Strings := calExpressionVarsFormatStringsByProto(expression)
	format := getExpressionFuncFormatByProto(expression)
	return fmt.Sprintf(format, vars2Strings...)
}

// calAppendTextByColumnAsByProto 构建 Column 类型的格式化字符串填充
func calAppendTextByColumnAsByProto(format string, column *pb.Column, userAs bool, alias string) string {
	if userAs {
		return fmt.Sprintf(format, column.Table, column.Field, alias)
	}
	return fmt.Sprintf(format, column.Table, column.Field)
}

// getExpressionFuncFormatByProto 获取Expression对应的格式化字符串
func getExpressionFuncFormatByProto(expression *pb.Expression) string {
	callType := expression.CallType
	// Expression 表达式只有在select关键字使用alias
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

// calExpressionVarsFormatStringsByProto 构建Expression表达式占位符所需的字符串填充
func calExpressionVarsFormatStringsByProto(expr *pb.Expression) []interface{} {
	vars := expr.Vars
	// 表达式参数: int, string, clause.Column
	var expressions []interface{}

	// 如果函数是 count1, 只用填充别名
	if expr.Call == string(clause.Count1) {
		expressions = append(expressions, expr.CallAs)
		return expressions
	}

	for _, varItem := range vars {
		switch v := varItem.GetVars().(type) {
		case *pb.MixVars_Column:
			column := v.Column
			format := util.Ternary(column.UseAs, clause.PGTableColumnAs, clause.PGTableColumn).(string)
			expressions = append(expressions, calAppendTextByColumnAsByProto(format, column, column.UseAs, expr.CallAs))
		case *pb.MixVars_Context:
			exprString := v.Context
			expressions = append(expressions, exprString)
		case *pb.MixVars_Number:
			exprInt := v.Number
			expressions = append(expressions, fmt.Sprintf(clause.AnyLiteral, exprInt))
		default:
			continue
		}
	}

	if expr.UseAs {
		expressions = append(expressions, expr.CallAs)
	}

	return expressions
}

func (f *PostgresModelBuilderFacade) BuildOther(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
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
	if sqlRef.Limit != nil {
		builder = buildLimitProto(builder, sqlRef)
	}

	return builder
}

func buildJoinByProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	return builder
}
func buildWhereByProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {

	return builder
}

func buildOrderByProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	orderByClause := sqlRef.OrderBy
	// 提取 orderBy 字符串片段
	var orderBys []string
	for _, orderBy := range orderByClause {
		field := processMixFieldItemByProto(orderBy.Dependent)
		orderByItem := fmt.Sprintf(clause.DoubleStringLiteral, field, orderBy.Order)
		orderBys = append(orderBys, orderByItem)
	}
	// 合并成一个字符串片段
	orderByItems := strings.Join(orderBys, clause.COMMA)
	// xorm 的 orderBy 方法只接收一个字符串参数，因此要自己将多个 orderBy 拼接为1个
	orderByFragment := fmt.Sprintf(clause.StringLiteral, orderByItems)

	return builder.OrderBy(orderByFragment)
}

func buildLimitProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	limit := sqlRef.Limit
	// 分页大小为0数据库允许但代码视为无效分页
	if limit.LimitN == 0 {
		return builder
	}
	return builder.Limit(int(limit.LimitN), int(limit.Offset))
}

func buildGroupByByProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	groupByClause := sqlRef.GroupBy
	// 类型断言: Column 或者 Expression
	groupByItems := extraColumnOrExpressionStringsByProto(groupByClause)

	// 使用 strings.Join 来合并切片元素，以逗号分隔
	joinedItems := strings.Join(groupByItems, clause.COMMA)

	// xorm 的 GroupBy 方法只接收一个字符串参数，因此要自己将多个 GroupBy 拼接为1个
	groupByFragment := fmt.Sprintf(clause.StringLiteral, joinedItems)

	return builder.GroupBy(groupByFragment)
}
