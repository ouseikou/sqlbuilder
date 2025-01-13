package clause

import "strings"

type CompareExpression string // > = < >= <= !=
type ArithExpression string   // + - * / %
type LogicExpression string   // and, or, not, between, in, not in, like, is not null,

const (
	Eq  CompareExpression = "="
	Neq CompareExpression = "!="
	Gt  CompareExpression = ">"
	Gte CompareExpression = ">="
	Lt  CompareExpression = "<"
	Lte CompareExpression = "<="
)

const (
	Add ArithExpression = "+"
	Sub ArithExpression = "-"
	Mul ArithExpression = "*"
	Div ArithExpression = "/"
	Mod ArithExpression = "%"
)

const (
	And       LogicExpression = " and "
	Or        LogicExpression = " or "
	Not       LogicExpression = " not "
	Between   LogicExpression = " between "
	In        LogicExpression = " in "
	NotIn     LogicExpression = " not in "
	Like      LogicExpression = " like "
	IsNotNull LogicExpression = " is not null "
)

const (
	COMMA = ", "
)

const (
	PntFormat     = `(%s)`
	PGAsFormat    = `%s as "%%s"`
	PGAsAppend    = ` as "%s"`
	MYSQLAsFormat = "%s as `%%s`"
)

type Operator struct {
	Op interface{} // 基于上面3种类型
}

const (
	ArithAddFormat   = `%s + %s`
	ArithAddAsFormat = `%s + %s as "%s"`

	ArithSubFormat   = `%s - %s`
	ArithSubAsFormat = `%s - %s as "%s"`

	ArithMulFormat   = `%s * %s`
	ArithMulAsFormat = `%s * %s as "%s"`

	ArithDivFormat   = `%s / %s`
	ArithDivAsFormat = `%s / %s as "%s"`

	ArithModFormat   = `%s %% %s`
	ArithModAsFormat = `%s %% %s as "%s"`
)

// ArithFormatMap 内置函数名称和SQL片段 映射
var (
	ArithFormatMap = map[ArithExpression]string{
		Add: ArithAddFormat,
		Sub: ArithSubFormat,
		Mul: ArithMulFormat,
		Div: ArithDivFormat,
		Mod: ArithModFormat,
	}

	ArithAsFormatMap = map[ArithExpression]string{
		Add: ArithAddAsFormat,
		Sub: ArithSubAsFormat,
		Mul: ArithMulAsFormat,
		Div: ArithDivAsFormat,
		Mod: ArithModAsFormat,
	}
)

// CalArithFormat 计算算术运算符的format
// 参数:
//   - call: 运算符号
//   - argsLen: 参数个数
//
// 返回值:
//   - 返回 算术运算符的format
func CalArithFormat(call string, argLen int) string {
	// strings.Join(formatParts, " "+ call + " ")
	var formatParts []string
	for i := 0; i < argLen; i++ {
		formatParts = append(formatParts, "%s")
	}

	return strings.Join(formatParts, " "+call+" ")
}

func CalArithFormatWithBuilder(call string, argLen int) string {
	var builder strings.Builder

	// 用一个循环逐个构建字符串部分
	for i := 0; i < argLen; i++ {
		if i > 0 {
			// 如果不是第一个元素，则先添加运算符
			builder.WriteString(" " + call + " ")
		}
		// 为每个参数添加"%s"
		builder.WriteString("%s")
	}

	return builder.String()
}
