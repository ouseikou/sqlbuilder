package clause

type InnerFunc string // count, sum, avg, max, min, distinct; 数据库的聚合函数不支持嵌套调用

const (
	CallDateTrunc InnerFunc = "date_trunc"
)

const (
	DateTruncFormat   = `date_trunc('%v', %s)`
	DateTruncAsFormat = `date_trunc('%v', %s) as "%s"`
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
