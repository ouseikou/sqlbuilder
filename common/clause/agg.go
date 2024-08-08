package clause

type AggFunc string // count, sum, avg, max, min, distinct; 数据库的聚合函数不支持嵌套调用

const (
	Count    AggFunc = "count"
	Sum      AggFunc = "sum"
	Avg      AggFunc = "avg"
	Max      AggFunc = "max"
	Min      AggFunc = "min"
	Distinct AggFunc = "distinct"
	// Stddev pg 总体标准差
	Stddev AggFunc = "stddev"
	// Variance pg 方差
	Variance AggFunc = "variance"
)

const (
	CountFormat    = `count(%s)`
	SumFormat      = `sum(%s)`
	AvgFormat      = `avg(%s)`
	MaxFormat      = `max(%s)`
	MinFormat      = `min(%s)`
	DistinctFormat = `distinct(%s)`
	StddevFormat   = `stddev(%s)`
	VarianceFormat = `variance(%s)`

	PGCountAsFormat    = `count(%s) as "%s"`
	PGSumAsFormat      = `sum(%s) as "%s"`
	PGAvgAsFormat      = `avg(%s) as "%s"`
	PGMaxAsFormat      = `max(%s) as "%s"`
	PGMinAsFormat      = `min(%s) as "%s"`
	PGDistinctAsFormat = `distinct(%s) as "%s"`
	PGStddevFormatAs   = `stddev(%s) as "%s"`
	PGVarianceFormatAs = `variance(%s) as "%s"`
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
		Stddev:   StddevFormat,
		Variance: VarianceFormat,
	}

	PGAggregationAsFuncFormatMap = map[AggFunc]string{
		Count:    PGCountAsFormat,
		Sum:      PGSumAsFormat,
		Avg:      PGAvgAsFormat,
		Max:      PGMaxAsFormat,
		Min:      PGMinAsFormat,
		Distinct: PGDistinctAsFormat,
		Stddev:   PGStddevFormatAs,
		Variance: PGVarianceFormatAs,
	}
)

type AggClause struct {
	// 聚合字段
	Aggregations []Expression `json:"aggregations"`
}
