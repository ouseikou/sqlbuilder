package util

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"

	"google.golang.org/protobuf/encoding/protojson"
)

// ProtoJsonOperate proto 序列化保留零值和枚举零值
var ProtoJsonOperate = protojson.MarshalOptions{
	EmitUnpopulated: true,
	// UseEnumNumbers:  true,
}

func Serialize(data interface{}) ([]byte, error) {
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}

func Deserialize(serialized []byte, target interface{}) error {
	err := json.Unmarshal(serialized, target)
	if err != nil {
		return err
	}
	return nil
}

func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// SourceCodeSubstringPath 获取执行时源码的绝对路径
func SourceCodeSubstringPath(regx string) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error retrieving caller information")
		return ""
	}

	fmt.Printf("Current file: %s\n", file)

	// 编译正则表达式
	re, err := regexp.Compile(regx)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return ""
	}

	matches := re.FindAllString(file, -1)
	if matches == nil {
		fmt.Println("No matches found")
		return ""
	}

	return matches[0]
}

// RemoveElements A切片移除存在于B切片元素
// 参数:
//   - sliceA: A切片
//   - sliceB: B切片
//
// 返回值:
//   - 返回: 新A切片
func RemoveElements[T comparable](sliceA []T, sliceB []T) []T {
	// 创建一个 map 来存储 sliceB 中的元素
	elementsToRemove := make(map[T]struct{})
	for _, elem := range sliceB {
		elementsToRemove[elem] = struct{}{}
	}

	// 创建一个新的切片来存储结果
	var result []T
	for _, elem := range sliceA {
		if _, found := elementsToRemove[elem]; !found {
			result = append(result, elem)
		}
	}

	return result
}
