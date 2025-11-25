package clause

import (
	"fmt"
	"strings"
	"testing"
)

var (
	// 新用户id和企业
	newSets = map[string]interface{}{
		"user_id":       666, // 新用户id要修改
		"enterprise_id": 4494,
	}
	// 旧用户id和企业
	oldWhere = map[string]interface{}{
		"user_id":       10697,
		"enterprise_id": 652,
	}

	schema = "holder_bi"
)

func TestBuildSimpleSql(t *testing.T) {
	info := SimpleUpdate{
		Driver: DriverMysql,
		Schema: schema,
		Table:  "api_ds",
		Sets:   newSets,
		Where:  oldWhere,
	}

	sql := info.BuildSimpleSql()
	fmt.Println("生成SQL: \r\n", sql)
}

func TestBuildBatchSimpleSql(t *testing.T) {
	tables := []string{
		"card",
		"chart_filter_config",
		"collection",
		"custom_sorting_config",
		"data_fill_submittal",
		"data_import",
		"dataset_bookmark",
		"dataset_drill_config",
		"dynamic_template_dataset",
		"export_task_info",
		"public_share",
		"report_view_base",
		"fill_field_map",
		"view_global_chart",
		"view_global_config",
		"view_global_filter",
		"view_global_table",
	}

	var sqlBatch []string
	for _, tableItem := range tables {
		item := SimpleUpdate{
			Driver: DriverMysql,
			Schema: schema,
			Table:  tableItem,
			Sets:   newSets,
			Where:  oldWhere,
		}
		sql := item.BuildSimpleSql()
		sqlBatch = append(sqlBatch, sql)
	}

	sqlBatchStr := strings.Join(sqlBatch, "\r\n")
	fmt.Println("生成SQL:\r\n", sqlBatchStr)
}
