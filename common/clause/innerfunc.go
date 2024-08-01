package clause

type InnerFunc string // count, sum, avg, max, min, distinct; 数据库的聚合函数不支持嵌套调用

const (
	CallDateTrunc InnerFunc = "DATE_TRUNC"
)

const (
	DateTruncFormat   = `DATE_TRUNC('%v', %s)`
	DateTruncAsFormat = `DATE_TRUNC('%v', %s) as "%s"`
)

// InnerFuncFormatMap 内置函数名称和SQL片段 映射
var (
	InnerFuncFormatMap = map[InnerFunc]string{
		CallDateTrunc: DateTruncFormat,
	}

	InnerAsFuncFormatMap = map[InnerFunc]string{
		CallDateTrunc: DateTruncAsFormat,
	}
)
