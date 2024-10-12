package facade

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

const (
	EmptyResultSql = "select * from (select 0 as x) src where src.x = 1"
)

const (
	TemplateDataStringFormat = `'%s'`
	TemplateDataNumberFormat = `%v`
	TemplateDataOtherFormat  = `'%v'`
)

func InjectFunc() template.FuncMap {
	return template.FuncMap{
		// MBQL用法

		// 1. 权限
		"PowerList": PowerList,

		// 新版模板语法内置函数

		// 规范化单参数
		"N1V": Normalize1Val,
		// 规范化数组
		"N1Arr": Normalize1Arr,
	}
}

func PowerList(guids []interface{}) string {
	var tempArr []string
	for i := range guids {
		val, ok := guids[i].(string)
		if ok {
			tempArr = append(tempArr, val)
		}
	}

	powerListStr := strings.Join(tempArr, ",")

	oneStr := fmt.Sprintf(TemplateDataStringFormat, powerListStr)

	return oneStr
}

func PowerListByString(guids []string) string {

	powerListStr := strings.Join(guids, ",")

	oneStr := fmt.Sprintf(TemplateDataStringFormat, powerListStr)

	return oneStr
}

// Normalize1Val 用法: {{N1V .Arg}} , 目的: 将1个参数值, 规范化为模板SQL的参数值
// 参数:
//   - arg: 1个参数值
//
// 返回值:
//   - 规范化(SQL)参数值
func Normalize1Val(arg interface{}) string {
	if nil == arg {
		return ""
	}
	switch arg := arg.(type) {
	case string:
		return fmt.Sprintf(TemplateDataStringFormat, arg)
	case float32, float64, int, int32, int64, bool:
		return fmt.Sprintf(TemplateDataNumberFormat, arg)
	default:
		return fmt.Sprintf(TemplateDataOtherFormat, arg)
	}
}

// Normalize1Arr 用法: {{N1Arr .Arr}} , 目的: 将1数组值, 规范化为模板SQL的参数值
// e.g. string : 'a','b','c' ; number/bool: a,b,c ; other: 'a','b','c'
// 参数:
//   - arr: 数组
//
// 返回值:
//   - 规范化(SQL)参数值
func Normalize1Arr(arr interface{}) string {

	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return ""
	}

	normalizeArr := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		normalizeArr[i] = Normalize1Val(item)
	}
	// string : 'a','b','c' ; num/bool: a,b,c ; other: 'a','b','c'
	return strings.Join(normalizeArr, ",")
}

// ----------------------------------------------------------------------------------------

func NormalizeValue(args map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range args {
		val := DataFormat(v)
		result[k] = val
	}
	// 返回规范化后的map
	return result
}

func DataFormat(v interface{}) interface{} {
	// 1. 对于 string类型的值v, 需要使用单引号包裹: `'%s'`
	// 2. 数值类型直接使用: `%v`
	// 3. 其他类型, : `'%v'`
	// 4. 数组按照上面的方式进行包裹
	switch arg := v.(type) {
	case string:
		return fmt.Sprintf(TemplateDataStringFormat, arg)
	case float32, float64, int, int32, int64, bool:
		return fmt.Sprintf(TemplateDataNumberFormat, arg)
	default:
		return fmt.Sprintf(TemplateDataOtherFormat, arg)
	}
}

func IsArray(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return true
	}

	return false
}
