package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"sql-builder/common/facade"
	"testing"
)

const TestText = `
		{{.Val1}}
		{{AFunc    .VB2  }}
		{{    .VB3    }}
		{{    .VB4_AA}} {{/* - 作为变量名会报错, 约束变量名: 英文字母下划线 */}}
		{{.VB4_AA1 }}
		{{BFunc (AFunc .VB1A .VB2B )}}
		.Val1
		{{range .Items}}{{.Name}}{{end}}  // .Items.Name
		{{with .Person}}{{.Name}}{{end}}
		{{if .PAYMENT_TYPE}} {{.NameC}}{{end}}
	`

// 通过正则提取, 存在问题: 1. 误伤 {{xxx "www.baidu.com"}}
func TestTemplateRegex(t *testing.T) {

	// 或者 \{\{(?:.*?){0,}(?<NAME>\.[a-z0-9]+)(?:.*?){0,}\}\} , 都会将嵌套自定义函数非首个参数丢掉,不清楚原因
	re := regexp.MustCompile(`{{[^}]*?(\.\w+)[^}]*}}`)
	matches := re.FindAllStringSubmatch(TestText, -1)

	for _, match := range matches {
		for i := 1; i < len(match); i++ {
			fmt.Println(match[i])
		}
	}
}

func TestExtraNodeVariable(test *testing.T) {

	data, err := os.ReadFile(filepath.Join("../assets/trans2.tmpl"))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := string(data)

	vars, _ := facade.ExtraValFromTemplate(content)
	sort.Strings(vars)

	for _, v := range vars {
		fmt.Println(v)
	}
}
