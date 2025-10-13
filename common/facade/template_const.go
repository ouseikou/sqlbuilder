package facade

import (
	"bytes"
	"github.com/ouseikou/sqlbuilder/common/clause"
	"text/template"
)

const (
	PGStartTimePlaceHolder = "first"
	PGEndTimePlaceHolder   = "second"

	// 对应 VisualColumnExpressionTypeEnum.INNER_PG_DATEDIFF_DAY
	PGDateDiffDayTmpl = `
        (DATE({{.second}}) - DATE({{.first}}))::INT
    `

	// 对应 VisualColumnExpressionTypeEnum.INNER_PG_DATEDIFF_MONTH
	PGDateDiffMonthTmpl = `
        (
          (EXTRACT(YEAR FROM {{.second}}) - EXTRACT(YEAR FROM {{.first}})) * 12
          + (EXTRACT(MONTH FROM {{.second}}) - EXTRACT(MONTH FROM {{.first}}))
          - CASE WHEN EXTRACT(DAY FROM {{.second}}) < EXTRACT(DAY FROM {{.first}}) THEN 1 ELSE 0 END
        )::INT
    `

	// 对应 VisualColumnExpressionTypeEnum.INNER_PG_DATEDIFF_YEAR
	PGDateDiffYearTmpl = `
        (
            EXTRACT(YEAR FROM {{.second}}) - EXTRACT(YEAR FROM {{.first}})
            - CASE
                WHEN (EXTRACT(MONTH FROM {{.second}}) < EXTRACT(MONTH FROM {{.first}}))
                  OR (EXTRACT(MONTH FROM {{.second}}) = EXTRACT(MONTH FROM {{.first}})
                      AND EXTRACT(DAY FROM {{.second}}) < EXTRACT(DAY FROM {{.first}}))
                THEN 1 ELSE 0
              END
          )::INT
    `
)

var (
	// 全量
	InnerFuncNameTemplateMap = map[string]string{
		string(clause.CallPgDateDiffDay):   PGDateDiffDayTmpl,
		string(clause.CallPgDateDiffMonth): PGDateDiffMonthTmpl,
		string(clause.CallPgDateDiffYear):  PGDateDiffYearTmpl,
	}

	InnerFuncPgDateDiffMap = map[string]string{
		string(clause.CallPgDateDiffDay):   PGDateDiffDayTmpl,
		string(clause.CallPgDateDiffMonth): PGDateDiffMonthTmpl,
		string(clause.CallPgDateDiffYear):  PGDateDiffYearTmpl,
	}
)

// RenderTmplLiteralInnerFunc 渲染Expression字面量表达式
func RenderTmplLiteralInnerFunc(goTmpl string, data map[string]string) (string, error) {
	t, err := template.New("sql").Parse(goTmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// IsPgDateDiffTmplFunc 是否是: pg_date_diff_?
func IsPgDateDiffTmplFunc(exp string) bool {
	return InnerFuncPgDateDiffMap[exp] != ""
}

// FormatTmplLiteralInnerFunc 解析Expression字面量表达式得到对应的完整字面量字符串
func FormatTmplLiteralInnerFunc(exp string, args []interface{}) string {
	if IsPgDateDiffTmplFunc(exp) {
		// 断言: args 的元素是string, 第一个参数是开始时间，第二个参数是结束时间, 转换为map[string]string
		// args []interface{} -> data map[string]string
		data := map[string]string{
			PGStartTimePlaceHolder: args[0].(string),
			PGEndTimePlaceHolder:   args[1].(string),
		}
		strLiteral, err := RenderTmplLiteralInnerFunc(InnerFuncNameTemplateMap[exp], data)
		if err != nil {
			panic(err)
		}
		return strLiteral
	}
	return ""
}
