package clause

import (
	"bytes"
	"fmt"
	"strings"
)

type SimpleUpdate struct {
	Driver Driver
	Schema string
	Table  string
	Sets   map[string]interface{}
	Where  map[string]interface{}
}

var (
	string2LiteralSafeMapOnJson = map[Driver]string{
		DriverPostgres: PGTableColumn,
		DriverMysql:    MysqlTableTableColumn,
		DriverDoris:    MysqlTableTableColumn,
	}
)

const (
	leftEqRightFormat = " %s = %s "
)

func (info *SimpleUpdate) BuildSimpleSql() string {
	strBuf := bytes.NewBufferString("")

	// 关键字1
	strBuf.WriteString("update")

	str2Format := string2LiteralSafeMapOnJson[info.Driver]
	strBuf.WriteString(fmt.Sprintf(str2Format, info.Schema, info.Table))

	// 关键字2
	strBuf.WriteString(" set ")

	// 构建 set
	var tempSetArr []string
	for field, value := range info.Sets {
		left := fmt.Sprintf(str2Format, info.Table, fmt.Sprintf("%v", field))
		tempSetArr = append(tempSetArr, fmt.Sprintf(leftEqRightFormat, left, fmt.Sprintf("%v", value)))
	}

	finalSetStr := strings.Join(tempSetArr, ",")
	strBuf.WriteString(finalSetStr)

	// 关键字3
	strBuf.WriteString(" where ")

	// 构建 where
	var tempWhereArr []string
	for field, value := range info.Where {
		left := fmt.Sprintf(str2Format, info.Table, fmt.Sprintf("%v", field))
		tempWhereArr = append(tempWhereArr, fmt.Sprintf(leftEqRightFormat, left, fmt.Sprintf("%v", value)))
	}

	finalWhereStr := strings.Join(tempWhereArr, ",")
	strBuf.WriteString(finalWhereStr)

	return strBuf.String()
}
