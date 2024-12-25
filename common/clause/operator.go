package clause

type compareExpression string // > = < >= <= !=
type ArithExpression string   // + - * / %
type LogicExpression string   // and, or, not, between, in, not in, like, is not null,

const (
	Eq  compareExpression = "="
	Neq compareExpression = "!="
	Gt  compareExpression = ">"
	Gte compareExpression = ">="
	Lt  compareExpression = "<"
	Lte compareExpression = "<="
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
