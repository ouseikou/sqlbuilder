package facade

import (
	"errors"
	"fmt"
	"strconv"
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

// 接口实现检查
var _ ModelBuilderFacade = &PostgresModelBuilderFacade{}

// BuildSQL postgres 对于门面接口的实现
// 1. 构建基础内容: selects, from
// 2. 构建其他内容: where, group by, having(嵌套子查询), order by, limit&offset
// 参数:
//   - request: proto 请求对象
//
// 返回值:
//   - 返回 sql
//   - 异常
func (f *PostgresModelBuilderFacade) BuildSQL(request *pb.BuilderRequest) (string, error) {
	builders := request.Builders
	var appendBuilder *xorm.Builder

	// 循环 append *xorm.Builder 内容
	for _, ref := range builders {
		xormBuilder, err := f.processBuilders(appendBuilder, ref)
		if err != nil {
			return "", err
		}
		appendBuilder = xormBuilder
	}

	if appendBuilder == nil {
		return "", errors.New("nest appendBuilder is nil")
	}

	// builder -> sql
	boundSQL, err := appendBuilder.ToBoundSQL()
	if err != nil {
		return "", err
	}

	return boundSQL, nil
}

// processBuilders 构建单层完整构建器
// 参数:
//   - appendBuilder: proto 请求对象
//   - builderRef: proto 请求中当前层级sqlRef对象
//
// 返回值:
//   - 返回 *xorm.Builder
//   - 异常
func (f *PostgresModelBuilderFacade) processBuilders(
	appendBuilder *xorm.Builder,
	builderRef *pb.DeepWrapper,
) (*xorm.Builder, error) {
	// 上游确定是 model 构造器
	sqlRef := builderRef.Sql.GetModel()
	// 选择方言
	dialectBuilder, err := f.BuildDialect()
	if err != nil {
		return nil, err
	}
	// builder(selects, from , as)
	basicBuilder := f.BuildBasic(appendBuilder, dialectBuilder, sqlRef)
	// 调用BuildOther方法处理其他逻辑
	builder := f.BuildOther(basicBuilder, sqlRef)

	return builder, nil
}

func (f *PostgresModelBuilderFacade) BuildDialect() (*xorm.Builder, error) {
	logger.Debug("[模板方法模式]-[postgres] 实现: 构建方言 builder...")
	return xorm.Dialect(xorm.POSTGRES), nil
}

func (f *PostgresModelBuilderFacade) BuildBasic(
	appendBuilder *xorm.Builder,
	dialectBuilder *xorm.Builder,
	sqlRef *pb.SqlReference,
) *xorm.Builder {
	logger.Debug("[模板方法模式]-[postgres] 实现: 构建 builder 基础内容...")

	// 嵌套sql, 构建 select + from
	if appendBuilder != nil {
		nestBasicBuilder := buildNestBasicByProto(appendBuilder, dialectBuilder, sqlRef)
		return nestBasicBuilder
	}
	// 一层sql, 构建 select + from
	selects, tableSchema, fromAlias := createBasicArgsByProto(sqlRef)

	basicBuilder := buildBasicByProto(dialectBuilder, selects, tableSchema, fromAlias)

	return basicBuilder
}

// buildNestBasicByProto 构建嵌套sql的 select + from
// 参数:
//   - appendBuilder: 上一次的 xorm.Builder 指针对象
//   - dialectBuilder: 当前层级 *xorm.Builder 初始化指针对象
//   - sqlRef: 当前层级的sqlRef对象
//
// 返回值:
//   - 返回 *xorm.Builder
func buildNestBasicByProto(
	appendBuilder *xorm.Builder,
	dialectBuilder *xorm.Builder,
	sqlRef *pb.SqlReference,
) *xorm.Builder {
	// 类型断言: Column 或者 Expression
	selects := extraColumnOrExpressionStringsByProto(sqlRef.GetSelect())

	nestBasicBuilder := dialectBuilder.
		Select(selects...).
		From(appendBuilder, fmt.Sprintf(clause.PGStringLiteralSafe, sqlRef.GetFrom().GetTableAlias()))

	return nestBasicBuilder
}

// buildBasicByProto 构建基本的SQL查询器
// 该函数旨在通过XORM库构建一个基本的SQL查询器对象
// 参数:
//   - dialectBuilder: XORM的Builder对象，用于构建SQL语句
//   - selects: 需要选择的列的列表
//   - sub: &schema.table/子查询的标识符
//   - fromAlias: &schema.table的表别名/FROM子句中的表别名
//
// 返回值:
//   - 返回一个构建好的SQL查询器对象
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
// 参数:
//   - sqlRef: 当前层级的sqlRef对象
//
// 返回值:
//   - 返回 .Select($slice)
//   - 返回 .From(&schema.table, $fromAlias)
func createBasicArgsByProto(sqlRef *pb.SqlReference) ([]string, string, string) {
	selectClause := sqlRef.Select

	// 类型断言: Column 或者 Expression
	selects := extraColumnOrExpressionStringsByProto(selectClause)

	fromClause := sqlRef.From

	tableSchema := fmt.Sprintf(clause.PGTableColumn, fromClause.TableSchema, fromClause.TableName)

	return selects, tableSchema, fmt.Sprintf(clause.PGStringLiteralSafe, fromClause.TableAlias)
}

// extraColumnOrExpressionStringsByProto 将MixField转换为对应的string字符串
// 参数:
//   - mixClause: proto 中的 MixField 对象
//
// 返回值:
//   - 返回 .Select($slice) 的 $slice 切片
func extraColumnOrExpressionStringsByProto(mixClause []*pb.MixField) []string {
	var selects []string
	for _, mixItem := range mixClause {
		result := processMixFieldItemByProto(mixItem)
		if result != "" {
			selects = append(selects, result)
		}
	}
	return selects
}

// processMixFieldItemByProto 将MixField转换为对应的string字符串
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
	case *pb.MixField_CaseWhen:
		caseWhen := v.CaseWhen
		return formatCaseWhenByProto(caseWhen)
	default:
		return ""
	}
}

// formatCaseWhenByProto 构建 case-when 类型的格式化字符串填充
// 参数:
//   - caseWhen: proto 中的 caseWhen 对象
//
// 返回值:
//   - 返回 caseWhen 对应字符串填充
func formatCaseWhenByProto(caseWhen *pb.CaseWhen) string {

	// proto -> struct(case-when) -> string
	caseWhenThens := formatCaseWhenItemByProto(caseWhen)

	caseWhenStruct := clause.CaseWhen{
		Conditions: caseWhenThens,
		// case-when 默认没有 else 分支
		ElseValue: "",
		Alias:     caseWhen.Alias,
	}

	fragment := caseWhenStruct.BuilderCaseWhenFragment(clause.DriverPostgres)
	// *xorm.Builder.Select(item...)  会自动填充逗号, 这里返回当前 items 切片的填充内容
	return fragment
}

func formatCaseWhenItemByProto(caseWhen *pb.CaseWhen) []clause.CaseWhenCondition {
	var conditions []clause.CaseWhenCondition

	for _, c := range caseWhen.GetConditions() {
		thenValue := ExtraArgItemValue(c.GetThen())
		whenValue := mixWhere2Condition(c.GetWhen())
		// 构造一行 case-when-then 语句
		whenThenLine := clause.CaseWhenCondition{
			Then: thenValue,
			When: whenValue,
		}

		conditions = append(conditions, whenThenLine)
	}

	return conditions
}

// formatColumnByProto 构建 Column 类型的格式化字符串填充
// 参数:
//   - column: proto 中的 Column 对象
//   - format: 格式化字符串
//
// 返回值:
//   - 返回 Column 对应字符串填充
func formatColumnByProto(column *pb.Column, format string) string {
	return calAppendTextByColumnAsByProto(format, column, column.UseAs, column.Alias)
}

// formatExpressionByProto 构建 Expression 类型的格式化字符串填充
func formatExpressionByProto(expression *pb.Expression) string {
	// 通常表达式的format中占位符和表达式参数是一一对应的拉链结构
	vars2Strings := calExpressionVarsFormatStringsByProto(expression)
	format := getExpressionFuncFormatByProto(expression)
	return fmt.Sprintf(format, vars2Strings...)
}

// calAppendTextByColumnAsByProto 构建 Column 类型的格式化字符串填充
// 参数:
//   - format: 格式化字符串
//   - column: proto 中的 Column 对象
//   - userAs: format 是否使用 as
//   - alias: format 字符串 as 占位符对应别名
//
// 返回值:
//   - 返回 Column 对应字符串填充
func calAppendTextByColumnAsByProto(format string, column *pb.Column, userAs bool, alias string) string {
	if userAs {
		return fmt.Sprintf(format, column.Table, column.Field, alias)
	}
	return fmt.Sprintf(format, column.Table, column.Field)
}

// getExpressionFuncFormatByProto 获取Expression对应的格式化字符串
// 参数:
//   - expression: proto 中的 Expression 对象
//
// 返回值:
//   - 返回 Expression 对应字符串填充
func getExpressionFuncFormatByProto(expression *pb.Expression) string {
	// 动态构建 Expression 的 format
	var finalFormat string

	// 初始化 Expression 的 format, 取不要 as 的那一部分
	basicFormat := expressionFuncNoAsFormat(expression)
	finalFormat = basicFormat

	// 是否使用小括号
	if expression.UsePnt {
		finalFormat = fmt.Sprintf(clause.PntFormat, finalFormat)
	}

	// 是否使用 as 别名
	// Expression 表达式只有在select关键字使用alias
	if expression.UseAs {
		finalFormat = fmt.Sprintf(clause.PGAsFormat, finalFormat)
	}

	return finalFormat
}

// expressionFuncNoAsFormat 获取Expression没有As的格式化字符串部分
// 参数:
//   - expression: proto 中的 Expression 对象
//
// 返回值:
//   - 返回 Expression 没有As的格式化字符串部分
func expressionFuncNoAsFormat(expression *pb.Expression) string {
	callType := expression.CallType
	switch callType {
	case pb.CallType_CALL_TYPE_AGG:
		return clause.AggregationFuncFormatMap[clause.AggFunc(expression.Call)]
	case pb.CallType_CALL_TYPE_INNER:
		return clause.InnerFuncFormatMap[clause.InnerFunc(expression.Call)]
	case pb.CallType_CALL_TYPE_ARITH:
		// 算术表达式的format要结合参数个数动态生成
		return clause.CalArithFormatWithBuilder(expression.Call, len(expression.Vars))
	default:
		return ""
	}
}

// calExpressionVarsFormatStringsByProto 构建Expression表达式占位符所需的字符串填充
// 参数:
//   - expression: proto 中的 Expression 对象
//
// 返回值:
//   - 返回 Expression 表达式参数的字符串切片
func calExpressionVarsFormatStringsByProto(expr *pb.Expression) []interface{} {
	vars := expr.Vars
	// 表达式参数: int, string, clause.Column
	var expressions []interface{}
	// var的字面量值 -> 是否转义, 只针对 Column 和 Expression 类型的 var
	var varTypeMap = make(map[any]bool)

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
			colLiteral := calAppendTextByColumnAsByProto(format, column, column.UseAs, expr.CallAs)
			varTypeMap[colLiteral] = false
			expressions = append(expressions, colLiteral)
		case *pb.MixVars_Context:
			exprString := v.Context
			varTypeMap[exprString] = true
			expressions = append(expressions, exprString)
		case *pb.MixVars_Number:
			exprInt := v.Number
			varNumLiteral := fmt.Sprintf(clause.AnyLiteral, exprInt)
			expressions = append(expressions, varNumLiteral)
		case *pb.MixVars_DoubleNum:
			exprDouble := v.DoubleNum
			doubleNumLiteral := fmt.Sprintf(clause.AnyLiteral, exprDouble)
			expressions = append(expressions, doubleNumLiteral)
		case *pb.MixVars_Expression:
			exprVarExpression := v.Expression
			varExpItemStr := formatExpressionByProto(exprVarExpression)
			varTypeMap[varExpItemStr] = false
			expressions = append(expressions, varExpItemStr)
		default:
			continue
		}
	}

	// 如果是可变参数函数，将参数转换为一个字符串
	if clause.IsVariadicArgsFunc(expr.Call) {
		expressions = []interface{}{variadicArgsFuncVars2oneVar(expressions, varTypeMap)}
	}

	if expr.UseAs {
		expressions = append(expressions, expr.CallAs)
	}

	return expressions
}

// variadicArgsFuncVars2oneVar 将可变参数函数的参数转换为一个字符串
// 参数:
//   - vars: proto 中的 Expression.vars
//
// 返回值:
//   - 合并字符串
func variadicArgsFuncVars2oneVar(vars []interface{}, varTypeMap map[any]bool) string {
	var builder strings.Builder
	// 遍历 args，将每个参数拼接成适合的 SQL 字符串
	for i, arg := range vars {
		// 对每个参数进行转换，确保它是一个有效的 SQL 字符串
		if i > 0 {
			builder.WriteString(", ")
		}

		switch v := arg.(type) {
		case string:
			// 如果是字符串类型原本是 Column 和 Expression, 不需要加单引号
			if varTypeMap[arg] == false {
				builder.WriteString(fmt.Sprintf("%s", v))
				break
			}
			// 字符串类型加上单引号
			builder.WriteString(fmt.Sprintf("'%s'", v))

		case int:
			// int 类型转换为字符串
			builder.WriteString(strconv.Itoa(v))
		case int32:
			// int32 类型转换为字符串
			builder.WriteString(strconv.FormatInt(int64(v), 10))
		case int64:
			// 数字类型直接转换为字符串
			builder.WriteString(strconv.FormatInt(v, 10))
		case float32, float64:
			// 浮点数类型直接转换为字符串
			builder.WriteString(fmt.Sprintf("%f", v))
		default:
			// 默认类型处理（其他类型，可以根据需要扩展）
			builder.WriteString(fmt.Sprintf("%v", arg))
		}
	}

	return builder.String()
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

// buildJoinByProto 根据SqlReference对象构建SQL查询。
// 参数:
//   - builder: *xorm.Builder类型的指针，用于构建SQL查询
//   - sqlRef: *pb.SqlReference类型的指针，包含SQL查询的引用信息
//
// 返回值:
//   - *xorm.Builder类型的指针，用于继续构建SQL查询
func buildJoinByProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	return builder
}

// buildWhereByProto 根据SqlReference构建WHERE条件
// 参数:
//   - builder: xorm.Builder的指针，用于构建SQL查询
//   - sqlRef: 指向SqlReference的指针，包含要构建WHERE条件的定义
//
// 返回:
//   - *xorm.Builder: 返回构建好的xorm.Builder，带有WHERE条件
func buildWhereByProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	xormCond := mixWhere2Condition(sqlRef.GetWhere())
	return builder.Where(xormCond)
}

// mixWhere2Condition 根据SqlReference构建WHERE条件
// 参数:
//   - wheres: 指向WHERE条件的指针
//
// 返回:
//   - xorm.Cond: 条件结构体 xorm.Builder.Cond
func mixWhere2Condition(wheres []*pb.MixWhere) xorm.Cond {
	var xormCond xorm.Cond

	for _, whereItem := range wheres {
		switch f := whereItem.GetFilter().(type) {
		case *pb.MixWhere_Condition:
			// 为什么 xormCond 每次经过forr都没有刷新值,而都是nil
			xormCond = buildWhereConditions(xormCond, f.Condition)
		case *pb.MixWhere_Expression:
			// 暂时不考虑 Where-Expression, e.g. WHERE ARRAY_CONTAINS(tags, 'urgent')
		}
	}
	return xormCond
}

// buildWhereConditions 构建WHERE条件
// 该方法主要用于根据给定的protobuf条件对象构建对应的xorm条件对象
// 它会考虑与现有条件的逻辑连接方式（AND或OR），并据此构建复合条件
// 参数:
//   - xormCond: 当前的xorm条件对象，可能为nil
//   - condition: protobuf定义的条件对象，用于构建新的条件
//
// 返回值:
//   - 返回更新后的xorm条件对象
func buildWhereConditions(xormCond xorm.Cond, condition *pb.Condition) xorm.Cond {
	currentCondTemp := buildWhereConditionItem(condition)
	// 通过 reverse 判断是否对当前条件取反
	currentCondition := refreshReverseWhereConditionItem(condition, currentCondTemp)
	if nil == xormCond {
		// 没有where条件, 作为第一个条件
		return currentCondition
	}
	// 处理当前条件和之前条件的逻辑连接; Where().And().Or(), 会将 当前cond片段 和 之前cond片段(左侧片段可能很长)分别视作两个不同整体
	if pb.Logic_LOGIC_AND == condition.Logic {
		return xorm.And(xormCond, currentCondition)
	}
	return xorm.Or(xormCond, currentCondition)
}

// refreshReverseWhereConditionItem 构建WHERE条件
// 通过 reverse 判断是否对当前条件取反
// 参数:
//   - pbCond: protobuf定义的条件对象，用于构建新的条件
//   - builderCond: 已经构建的新条件
//
// 返回值:
//   - 返回更新后的xorm条件对象
func refreshReverseWhereConditionItem(pbCond *pb.Condition, builderCond xorm.Cond) xorm.Cond {
	if pbCond.GetReverse() {
		return xorm.Not{builderCond}
	}
	return builderCond
}

// buildWhereConditionItem 根据提供的Condition对象构建xorm的条件表达式
// 主要用于数据库查询时的条件构造，通过解析Condition中的操作符和参数，生成对应的xorm条件对象
// 参数:
//   - cond 指向一个protobuf定义的Condition对象，用于构建查询条件
//
// 返回值:
//   - 返回生成的xorm.Cond对象，用于数据库查询；如果无法构建，则返回nil
func buildWhereConditionItem(cond *pb.Condition) xorm.Cond {
	operator := cond.Operator
	switch operator {
	//column op arg[0]
	case pb.Op_OP_EQ:
		field := extraConditionOpField(cond)
		return &xorm.Eq{field: ExtraArgItemValue(cond.Args[0])}
	case pb.Op_OP_NEQ:
		field := extraConditionOpField(cond)
		return &xorm.Neq{field: ExtraArgItemValue(cond.Args[0])}
	case pb.Op_OP_LT:
		field := extraConditionOpField(cond)
		return &xorm.Lt{field: ExtraArgItemValue(cond.Args[0])}
	case pb.Op_OP_LTE:
		field := extraConditionOpField(cond)
		return xorm.Lte{field: ExtraArgItemValue(cond.Args[0])}
	case pb.Op_OP_GT:
		field := extraConditionOpField(cond)
		return xorm.Gt{field: ExtraArgItemValue(cond.Args[0])}
	case pb.Op_OP_GTE:
		field := extraConditionOpField(cond)
		return xorm.Gte{field: ExtraArgItemValue(cond.Args[0])}
	case pb.Op_OP_LIKE:
		field := extraConditionOpField(cond)
		return xorm.Like{field, ExtraArgItemValue(cond.Args[0]).(string)}
		// column op ...arg
	case pb.Op_OP_BETWEEN:
		field := extraConditionOpField(cond)
		return xorm.Between{Col: field, LessVal: ExtraArgItemValue(cond.Args[0]), MoreVal: ExtraArgItemValue(cond.Args[1])}
	case pb.Op_OP_IN:
		field := extraConditionOpField(cond)
		inArgs := extraConditionArgsSlice(cond)
		return xorm.In(field, inArgs)
	case pb.Op_OP_NOT_IN:
		field := extraConditionOpField(cond)
		inArgs := extraConditionArgsSlice(cond)
		return xorm.NotIn(field, inArgs)
		// column op
	case pb.Op_OP_IS_NULL:
		field := extraConditionOpField(cond)
		return xorm.IsNull{field}
	case pb.Op_OP_IS_NOT_NULL:
		field := extraConditionOpField(cond)
		return xorm.NotNull{field}
		// Where(xorm.Expr(Column, ...arg)), and 和 or 暂时使用Expr处理, 需要sqlbuilder调用方提供完整 sqlFormat;; => 未测试
	case pb.Op_OP_AND:
		formatField := extraConditionOpField(cond)
		formatArgs := extraConditionArgsSlice(cond)
		return xorm.And(xorm.Expr(formatField, formatArgs))
	case pb.Op_OP_OR:
		formatField := extraConditionOpField(cond)
		formatArgs := extraConditionArgsSlice(cond)
		return xorm.Or(xorm.Expr(formatField, formatArgs))
	}
	return nil
}

// extraConditionOpField 根据给定的条件生成额外的操作字段字符串
// 这个函数主要处理条件中的字段，通过调用processMixFieldItemByProto函数来实现字段的处理
// 参数:
//   - cond 指向一个protobuf定义的Condition对象，该对象包含了查询或操作的条件
//
// 返回值:
//   - 返回一个字符串，该字符串是处理后的结果
func extraConditionOpField(cond *pb.Condition) string {
	return processMixFieldItemByProto(cond.GetField())
}

// ExtraArgItemValue 根据BasicData中的不同类型返回相应的值。
// 参数:
//   - item: 指向BasicData的指针作为参数，并根据BasicData中封装的实际数据类型
//
// 返回值:
//   - 返回一个接口类型值，可以是字符串、整数、浮点数或布尔值
func ExtraArgItemValue(item *pb.BasicData) interface{} {
	switch elem := item.GetData().(type) {
	case *pb.BasicData_StrVal:
		return elem.StrVal
	case *pb.BasicData_IntVal:
		return elem.IntVal
	case *pb.BasicData_DoubleVal:
		return elem.DoubleVal
	case *pb.BasicData_Logic:
		return elem.Logic
	}
	return ""
}

// extraConditionArgsSlice 根据给定的Condition对象不定长参数的切片
// 参数:
//   - cond 指向一个protobuf定义的Condition对象，该对象包含了查询或操作的条件
//
// 返回值:
//   - 返回一个接口类型值，可以是字符串、整数、浮点数或布尔值
func extraConditionArgsSlice(cond *pb.Condition) interface{} {
	var inArgs []interface{}
	for _, data := range cond.GetArgs() {
		switch d := data.GetData().(type) {
		case *pb.BasicData_IntVal:
			inArgs = append(inArgs, d.IntVal)
		case *pb.BasicData_DoubleVal:
			inArgs = append(inArgs, d.DoubleVal)
		case *pb.BasicData_StrVal:
			inArgs = append(inArgs, d.StrVal)
		case *pb.BasicData_Logic:
			inArgs = append(inArgs, d.Logic)
		}
	}
	return inArgs
}

// buildOrderByProto 根据给定的SqlReference对象构建xorm的OrderBy方法
// 参数:
//   - builder: *xorm.Builder类型的指针，用于构建SQL查询
//   - sqlRef: *pb.SqlReference类型的指针，包含SQL查询的引用信息
//
// 返回值:
//   - *xorm.Builder类型的指针，用于继续构建SQL查询
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

// buildLimitProto 根据给定的SqlReference对象构建xorm的Limit方法
// 参数:
//   - builder: *xorm.Builder类型的指针，用于构建SQL查询
//   - sqlRef: *pb.SqlReference类型的指针，包含SQL查询的引用信息
//
// 返回值:
//   - *xorm.Builder类型的指针，用于继续构建SQL查询
func buildLimitProto(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	limit := sqlRef.Limit
	// 分页大小为0数据库允许但代码视为无效分页
	if limit.LimitN == 0 {
		return builder
	}
	return builder.Limit(int(limit.LimitN), int(limit.Offset))
}

// buildGroupByByProto 根据给定的SqlReference对象构建xorm的GroupBy方法
// 参数:
//   - builder: *xorm.Builder类型的指针，用于构建SQL查询
//   - sqlRef: *pb.SqlReference类型的指针，包含SQL查询的引用信息
//
// 返回值:
//   - *xorm.Builder类型的指针，用于继续构建SQL查询
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
