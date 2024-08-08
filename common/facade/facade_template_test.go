package facade

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func TestTemplateString(t *testing.T) {
	strA := `"%s"."%s" as temp`

	// 填充两个占位符, 返回填充后的字符串
	finalStr := fmt.Sprintf(strA, "a", "b")

	expected := `"a"."b" as temp`

	if finalStr != expected {
		t.Errorf("Expected %s, but got %s", expected, finalStr)
	}
}

type Inventory struct {
	Material    string
	Count       uint
	TestStrArrs []string
	TestIntArrs []int
	Empty       string
}

const templateString = `
	{{.Count}} of {{.Material}}
	{{if eq .Count  17 }} T1 {{else}} T0 {{end}}
	{{if gt .Count  17}} T1 {{else if le .Count  17}} T0 {{end}}
	{{"\"output\""}}
	{{len .Material}}
	{{index .Material 0 }} {{/* 返回索引指向值, 字符串返回字符ASCII值 */}}
	{{Greet .Material}}

	{{"测试range" | printf "%q"}}
	{{range $index, $element := .TestStrArrs}}
	  Element {{$index}}: {{$element}}
	{{else}}
	  No elements
	{{end}}

	{{"测试注入结构体" | printf "%q"}} {{/* %q 字符串字面值，必要时会采用安全的转义表示 */}}
	{{.}} {{/* .代表Execute注入内容 */}}

	{{"测试局部变量" | printf "%q"}}
	   {{- $total := 0 -}}  
	{{- range .TestIntArrs -}}
	  {{- $total = Add $total . -}} {{/* . 这里代表循环item,在最外层和局部代表的内容是不一样的; 大致理解.代表注入模板或局部的整个结构体 */}}
	{{- end }} {{/* - 代表前或后的空白和换行符 */}}

	Total: {{$total}}

	{{"模板套模板" | printf "%q"}}
	{{define "header"}} 
	Header: {{.}} {{/* . 代表传入名为header的template的入参值  */}}
	{{end -}}
	
	{{define "body"}} 
	Body: {{.}} 
	{{end -}}
	
	{{define "footer"}} 
	Footer: {{.}} 
	{{end -}}
	
	{{- template "header" .Material}}
	{{- template "body" .TestStrArrs}}
	{{- template "footer" .TestIntArrs}}

	{{"存在判断" | printf "%q"}}
	{{with .Empty}} T1 {{else}} T0 {{end}} {{/* with判断是否不为nil */}}
`

const aaa = "{{`\"output\"`}}"

// 标准库内没有的函数需要自定义,并且在指定模板上FuncMap中注册
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func add(a, b int) int {
	return a + b
}

func TestTemplate(t *testing.T) {
	testStrArrs := []string{"apple", "banana", "cherry"}
	testIntArrs := []int{1, 2, 3}
	sweaters := Inventory{
		Material:    "wool",
		Count:       17,
		TestStrArrs: testStrArrs,
		TestIntArrs: testIntArrs,
	}

	// 注册函数
	tmpl, err := template.New("测试模板").Funcs(template.FuncMap{
		"Greet": greet,
		"Add":   add,
	}).Parse(templateString)
	if err != nil {
		t.Fatalf("Error parsing template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, sweaters)
	if err != nil {
		t.Fatalf("Error executing template: %v", err)
	}

	fmt.Printf("%s\n", buf.String())
}

const sqlTemplate = `
	SELECT 
		team.username 姓名,
		userdetail.user_name, 
		team.user_identification account, 
		team.tiname 部门,
		team.triname 岗位, 
		userdetail.project_name 项目, 
		userdetail.types 类型,
		CASE 
			WHEN userdetail.severities = '-' THEN '' 
			ELSE userdetail.severities 
		END 严重程度, 
		userdetail.num 数量
	FROM (
		SELECT
			id,
			account,
			user_name,
			project_name,
			types,
			severities,
			num 
		FROM "sonar".tb_user_detailquality detail
		WHERE 
			severities IN ('LOW', 'MEDIUM', 'HIGH')
			{{- if .ProjectName }}
			AND project_name = ANY({{.ProjectName}}::text[])
			{{- end }}
			{{- if .TypeCode }}
			AND types = ANY({{.TypeCode}}::text[])
			{{- end }}
			{{- if .SeveritiesType }}
			AND severities = ANY({{.SeveritiesType}}::text[])
			{{- end }}
			{{- if .StartTime }}
			AND detail.create_time = {{.StartTime}}
			{{- end }}
	) userdetail
	INNER JOIN sonar.tb_mapping_user_structure mapuser 
		ON mapuser.external_part_user_id = userdetail.user_name
		AND mapuser.system_type = 1 
		AND mapuser.use_state = 1 
	RIGHT JOIN (
		SELECT 
			u.username,
			tui.user_identification,
			ti.name AS tiname,
			tri.name AS triname
		FROM 
			team.team_user_info tui
			JOIN team.team_role_info tri ON tui.main_post_id = tri.id  
			JOIN team.USER u ON tui.user_identification = u.account 
			JOIN team.team_info ti ON tri.team_info_id = ti.id
		WHERE 
			tui.team_info_id = 244
			{{- if .Husername }}
			AND u.username = ANY({{.Husername}}::text[])
			{{- end }}
	) team 
		ON mapuser.holder_id = team.user_identification
	ORDER BY 
		2, 4
`

type SQLParams struct {
	ProjectName    string
	TypeCode       string
	SeveritiesType string
	StartTime      string
	Husername      string
}

func TestTemplateDashboard2330(t *testing.T) {
	params := SQLParams{
		//ProjectName:    `{"Project1", "Project2"}`,
		//TypeCode:       `{"Type1", "Type2"}`,
		SeveritiesType: `{MEDIUM}`,
		StartTime:      `'2024-04-01'`,
		Husername:      `{王成皓}`,
	}

	tmpl, err := template.New("sql").Parse(sqlTemplate)
	if err != nil {
		t.Fatalf("Error parsing template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, params)
	if err != nil {
		t.Fatalf("Error executing template: %v", err)
	}

	fmt.Printf("%s\n", buf.String())
}

var (
	_ CCC = (*AAA)(nil)
	_ CCC = (*BBB)(nil)
)

type CCC interface {
	do()
}

type AAA struct {
}

type BBB struct {
}

func (a *AAA) do() {
	fmt.Println("aaa")
}

func (b *BBB) do() {

	fmt.Println("bbbb")
}

func TestTemplateDashboard230(t *testing.T) {
	var a CCC
	a = &BBB{}
	switch a.(type) {
	case *BBB:
		fmt.Printf("bbb")
	case *AAA:
		fmt.Println("bbb")
	}
}
