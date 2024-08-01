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

type Operator struct {
	Op interface{} // 基于上面3种类型
}
