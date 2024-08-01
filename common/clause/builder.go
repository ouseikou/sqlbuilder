package clause

import (
	"github.com/ouseikou/sqlbuilder/common"
)

type DeepWrapper struct {
	// 深度: 控制层级; e.g. 0, 代表最里层, 1~n, 代表依次包裹 0 层
	Deep int `json:"deep"`
	// 某个层级的SQL引用对象
	Sql SQLReference `json:"sql"`
}

type BuilderRequest struct {
	// 不同层级嵌套构建SQL
	SQLBuilders []DeepWrapper          `json:"builders"`
	Driver      common.Driver          `json:"driver"`
	Strategy    common.BuilderStrategy `json:"strategy"`
}

// SQLReference todo Select缺少自定义字段
type SQLReference struct {
	From  Table        `json:"from"`
	Join  []Join       `json:"join"`
	Where []Expression `json:"where"`
	// 类型断言: Column 或者 Expression
	GroupBy []interface{} `json:"groupBy"`
	// Aggregation 会加入 Select 片段
	Aggregation []Expression `json:"aggregation"`
	// 类型断言: Column 或者 Expression
	Select  []interface{} `json:"select"`
	OrderBy []OrderBy     `json:"orderBy"`
	Limit   LimitClause   `json:"limit"`
}

//type MixColumnAndExpression interface {
//	mix()
//}
//
//var (
//	_ MixColumnAndExpression = (*Column)(nil)
//	_ MixColumnAndExpression = (*Expression)(nil)
//)
