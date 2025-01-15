package clause

import (
	"fmt"
	"strings"

	xorm "xorm.io/builder"
)

// CaseWhen 用于动态构建 CASE WHEN 表达式, case 列 when 相当于 case when 等值条件
type CaseWhen struct {
	Conditions []CaseWhenCondition
	ElseValue  string
	// 表示 CASE 的列名as
	Alias string
	UseAs bool
}

// CaseWhenCondition 表示1行 case-when-then 语句
type CaseWhenCondition struct {
	When xorm.Cond   // 表示条件的表达式
	Then interface{} // 条件满足时返回的值
}

// BuilderCaseWhenFragment 实现 case when 片段
func (cw *CaseWhen) BuilderCaseWhenFragment(driver Driver) string {
	var builder strings.Builder
	// 开始构建 CASE WHEN 语句
	builder.WriteString("CASE ")

	// 构建 WHEN 条件
	for _, cond := range cw.Conditions {
		fragment, partErr := xorm.ToBoundSQL(cond.When)
		if partErr != nil {
			fmt.Println("Error building case when condition:", partErr)
			return ""
		}
		caseWhenSQL := fmt.Sprintf(" WHEN %s THEN '%v'", fragment, cond.Then)
		builder.WriteString(caseWhenSQL)
	}

	// 添加 ELSE 部分
	if cw.ElseValue != "" {
		caseWhenSQL := fmt.Sprintf(" ELSE '%s'", cw.ElseValue)
		builder.WriteString(caseWhenSQL)
	}

	// 结束 CASE 语句
	builder.WriteString(" END ")
	// 根据驱动返回转义符
	driverKeyword := DriverKeywordWrapMap[driver]
	if cw.UseAs {
		builder.WriteString(fmt.Sprintf(` as %s%s%s`, driverKeyword, cw.Alias, driverKeyword))
	}
	return builder.String()
}
