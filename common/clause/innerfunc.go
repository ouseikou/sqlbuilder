package clause

type InnerFunc string // count, sum, avg, max, min, distinct; 数据库的聚合函数不支持嵌套调用

const (
	CallDateTrunc InnerFunc = "date_trunc"
	CallPGToChar  InnerFunc = "to_char"
	CallRound     InnerFunc = "round"
	CallCeiling   InnerFunc = "ceiling"
	CallFloor     InnerFunc = "floor"
)

const (
	// DateTruncFormat DateTruncAsFormat 转换为离散时间, 最终返回是timestamp类型
	DateTruncFormat   = `date_trunc('%v', %s)`
	DateTruncAsFormat = `date_trunc('%v', %s) as "%s"`

	// ToCharFormat ToCharAsFormat to_char(时间类型字段, format), 最终返回是字符串
	ToCharFormat   = `to_char(%s, '%v')`
	ToCharAsFormat = `to_char(%s, '%v') as "%s"`

	RoundFormat   = `round(%s, %s)`
	RoundFormatAs = `round(%s, %s) as "%s"`

	CeilingFormat   = `ceiling(%s)`
	CeilingFormatAs = `ceiling(%s) as "%s"`

	FloorFormat   = `floor(%s)`
	FloorFormatAs = `floor(%s) as "%s"`
)

// InnerFuncFormatMap 内置函数名称和SQL片段 映射
var (
	InnerFuncFormatMap = map[InnerFunc]string{
		CallDateTrunc: DateTruncFormat,
		CallPGToChar:  ToCharFormat,
		CallRound:     RoundFormat,
		CallCeiling:   CeilingFormat,
		CallFloor:     FloorFormat,
	}

	InnerAsFuncFormatMap = map[InnerFunc]string{
		CallDateTrunc: DateTruncAsFormat,
		CallPGToChar:  ToCharAsFormat,
		CallRound:     RoundFormatAs,
		CallCeiling:   CeilingFormatAs,
		CallFloor:     FloorFormatAs,
	}
)
