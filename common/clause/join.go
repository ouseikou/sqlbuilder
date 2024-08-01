package clause

type JoinType string //	left join, right join, inner join, full join, cross join

const (
	LeftJoin  JoinType = "left join"
	RightJoin JoinType = "right join"
	InnerJoin JoinType = "inner join"
	FullJoin  JoinType = "full join"
	CrossJoin JoinType = "cross join"
)

type Join struct {
	// 连接方式
	Type JoinType `json:"type"`
	// 连接表
	Table Table `json:"table"`
	// 类型断言: FieldClause, 默认主表字段
	Left interface{} `json:"left"`
	// 类型断言: FieldClause, 默认从表字段
	Right interface{} `json:"right"`
	// 连接条件
	On LogicExpression `json:"on"`
}

type JoinClause struct {
	Join []Join
}
