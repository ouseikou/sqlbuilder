package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
	xorm "xorm.io/builder"
)

func TestXORMSQLBuilderUnion(t *testing.T) {
	builder := xorm.Select("*").From("a").Where(xorm.Eq{"status": "1"}).
		Union("all", xorm.Select("*").From("a").Where(xorm.Eq{"status": "2"})).
		Union("distinct", xorm.Select("*").From("a").Where(xorm.Eq{"status": "3"})).
		Union("", xorm.Select("*").From("a").Where(xorm.Eq{"status": "4"}))

	sql, args, err := builder.ToSQL()
	/**
	(SELECT * FROM a WHERE status=?)
	UNION ALL
	(SELECT * FROM a WHERE status=?)
	UNION DISTINCT
	(SELECT * FROM a WHERE status=?)
	UNION
	(SELECT * FROM a WHERE status=?)

	[1 2 3 4] <nil>
	*/
	t.Log(sql, args, err)

	// (SELECT * FROM a WHERE status='1') UNION ALL (SELECT * FROM a WHERE status='2') UNION DISTINCT (SELECT * FROM a WHERE status='3') UNION (SELECT * FROM a WHERE status='4') <nil>
	boundSQL, err := builder.ToBoundSQL()
	t.Log(boundSQL, err)

}

func TestXORMSQLBuilderSub(t *testing.T) {
	builder := xorm.Select("sub.id").From(xorm.Select("c").From("table1").Where(xorm.Eq{"a": 1}), "sub").Where(xorm.Eq{"b": 1})

	sql, args, err := builder.ToSQL()
	// SELECT sub.id FROM (SELECT c FROM table1 WHERE a=?) sub WHERE b=? [1 1] <nil>
	t.Log(sql, args, err)

	// SELECT sub.id FROM (SELECT c FROM table1 WHERE a=1) sub WHERE b=1 <nil>
	boundSQL, err := builder.ToBoundSQL()
	t.Log(boundSQL, err)

}

func TestXORMSQLBuilder(t *testing.T) {
	//builder := xorm.Dialect(xorm.POSTGRES).
	//	Select(`DATE_TRUNC('month', "t1"."created_at") as created_at`, `"t1"."category"`, `sum("t1"."price")`).
	//	From(`"sample_data"."products"`, "t1").
	//	GroupBy(`DATE_TRUNC('month', "t1"."created_at"), "t1"."category"`).
	//	OrderBy(`DATE_TRUNC('month', "t1"."created_at") asc, "t1"."category" desc`)
	//
	//boundSQL, _ := builder.ToBoundSQL()
	//t.Log(boundSQL)

	// SELECT c FROM table1 ORDER BY CASE WHEN owner_name LIKE 'a' THEN 0 ELSE 1 END
	//sqlB, _ := xorm.Select("c").From("table1").OrderBy(xorm.Expr("CASE WHEN owner_name LIKE ? THEN 0 ELSE 1 END", "a")).ToBoundSQL()
	//fmt.Println(sqlB)

	builderWhere := xorm.Dialect(xorm.POSTGRES).
		Select(`"t1"."category"`).
		From(`"sample_data"."products"`, "t1").
		//Where(xorm.Or(xorm.Eq{"c": 3, "a": "asda"}, xorm.Eq{"d": 4}))
		//Where(xorm.Expr("c >1 and d = 6 or (q between 1 and 3)")) // expr 允许一切原生写法
		//Where(xorm.Expr("c >1 and d = ? or (q between 1 and 3)", "asd")) // expr 还有 fmt.Sprintf() 效果
		//Where(xorm.In("col1", []int{1, 2, 10})) // kv中v是数组
		//Where(xorm.Between{Col: "col1", LessVal: 1, MoreVal: 10}) // betweenObj
		//Where(xorm.Eq{"id": 1}).And(xorm.Eq{"aaa": 1}).Or(xorm.Lte{"qqq": 100}).And(xorm.Eq{"ddd": "qwe"}) // and/or 会将当前视作右和前面所有视为左整体圆括号圈定
		//Where(xorm.Eq{"id": 1}).And(xorm.Eq{"aaa": 1}).And(xorm.Lte{"qqq": 100}).And(xorm.Eq{"ddd": "qwe"}) //所有都是and不会加圆括号
		//Where(xorm.Eq{"id": 1, "category": "Doohickey", "price": 10}) // 按照key排序, 且共用一个操作符
		//Where(xorm.Eq{`"product"."id"`: 1}).Where(xorm.Eq{`"product"."category"`: "Doohickey"}) // 默认将多个where按照and粘合
		//And(xorm.Eq{`"product"."id"`: 1}).Or(xorm.Eq{`"product"."category"`: "Doohickey"}) // 可以不申明where, 会自动在前面拼接where关键字
		Where(xorm.And(xorm.Eq{`"product"."id"`: 1}).Or(xorm.Eq{`"product"."category"`: "Doohickey"})) // 在第一个不申明条件前拼接where关键字不影响

	whereBoundSQL, _ := builderWhere.ToBoundSQL()
	t.Log(whereBoundSQL)

	//// 使用 expr 实现下面效果
	//// (a=1 AND b LIKE '%c%') OR (a=2 AND b LIKE '%g%')
	//segment := xorm.Eq{"a": 1}.And(xorm.Like{"b", "c"}).Or(xorm.Eq{"a": 2}.And(xorm.Like{"b", "g"})).
	//	And(xorm.Expr("c >? or (q between 1 and 3)", 10))
	//segmentBoundSQL, _ := xorm.ToBoundSQL(segment)
	//t.Log(segmentBoundSQL)
}

// result := fmt.Sprintf(template, name, accountNumber)

func TestExprCond(t *testing.T) {
	// 两个where之间默认填充 AND, 如果希望一个字符串包含多个表达式使用 Expr, 内置表达式使用对应 type

	b := xorm.Select("id").
		From("table1").
		Where(xorm.Expr("a=? OR b=?", 1, 2)).
		Where(xorm.Or(xorm.Eq{"c": 3}, xorm.Eq{"d": 4}))
	sql, args, err := b.ToSQL()
	assert.NoError(t, err)
	assert.EqualValues(t, "table1", b.TableName())
	assert.EqualValues(t, "SELECT id FROM table1 WHERE (a=? OR b=?) AND (c=? OR d=?)", sql)
	assert.EqualValues(t, []interface{}{1, 2, 3, 4}, args)

	// 下面示例只能结合 Expression 和 Condition 表达:
	// Expression.Vars{Condition无next, Condition有1next}
	logicNest := xorm.Or(xorm.Eq{"a": 3}, xorm.And(xorm.Eq{"b": 1}, xorm.Eq{"c": 2}))
	logicSql, err := xorm.ToBoundSQL(logicNest)
	assert.EqualValues(t, "a=3 OR (b=1 AND c=2)", logicSql)

}
