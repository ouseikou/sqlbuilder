package clause

import (
	"fmt"
	"testing"
)

func TestBuildSimpleSql(t *testing.T) {
	info := SimpleUpdate{
		Driver: DriverMysql,
		Schema: "holder_bi",
		Table:  "api_ds",
		// 新用户id和企业
		Sets: map[string]interface{}{
			"user_id":       666, // 新用户id要修改
			"enterprise_id": 4494,
		},
		// 旧用户id和企业
		Where: map[string]interface{}{
			"user_id":       10697,
			"enterprise_id": 652,
		},
	}

	sql := info.BuildSimpleSql()
	fmt.Println("生成SQL: \r\n", sql)
}
