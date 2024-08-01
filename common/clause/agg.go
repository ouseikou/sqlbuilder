package clause

type AggFunc string // count, sum, avg, max, min, distinct; 数据库的聚合函数不支持嵌套调用

const (
	Count    AggFunc = "count"
	Sum      AggFunc = "sum"
	Avg      AggFunc = "avg"
	Max      AggFunc = "max"
	Min      AggFunc = "min"
	Distinct AggFunc = "distinct"

	//CountAs    AggFunc = "countAs"
	//SumAs      AggFunc = "sumAs"
	//AvgAs      AggFunc = "avgAs"
	//MaxAs      AggFunc = "maxAs"
	//MinAs      AggFunc = "minAs"
	//DistinctAs AggFunc = "distinctAs"
)

const (
	CountFormat      = `count(%s)`
	CountAsFormat    = `count(%s) as "%s"`
	SumFormat        = `sum(%s)`
	SumAsFormat      = `sum(%s) as "%s"`
	AvgFormat        = `avg(%s)`
	AvgAsFormat      = `avg(%s) as "%s"`
	MaxFormat        = `max(%s)`
	MaxAsFormat      = `max(%s) as "%s"`
	MinFormat        = `min(%s)`
	MinAsFormat      = `min(%s) as "%s"`
	DistinctFormat   = `distinct(%s)`
	DistinctAsFormat = `distinct(%s) as "%s"`
)

// AggregationFuncFormatMap 映射聚合函数
var (
	AggregationFuncFormatMap = map[AggFunc]string{
		Count:    CountFormat,
		Sum:      SumFormat,
		Avg:      AvgFormat,
		Max:      MaxFormat,
		Min:      MinFormat,
		Distinct: DistinctFormat,
	}

	AggregationAsFuncFormatMap = map[AggFunc]string{
		Count:    CountAsFormat,
		Sum:      SumAsFormat,
		Avg:      AvgAsFormat,
		Max:      MaxAsFormat,
		Min:      MinAsFormat,
		Distinct: DistinctAsFormat,
	}
)

type AggClause struct {
	// 聚合字段
	Aggregations []Expression `json:"aggregations"`
}
