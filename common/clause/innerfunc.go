package clause

type InnerFunc string // count, sum, avg, max, min, distinct; 数据库的聚合函数不支持嵌套调用

// 无参数函数
const (
	Count1Str string = "count1"
	NowFunc          = "now"
)

// 无参数函数format
const (
	NowFormat = "now()"
)

// 通用函数
const (
	CallRound   InnerFunc = "round"
	CallCeiling InnerFunc = "ceiling"
	CallFloor   InnerFunc = "floor"
	Abs         InnerFunc = "abs"
	ModFunc     InnerFunc = "mod"
	Sqrt        InnerFunc = "sqrt"

	// CastType 类型强转使用表达式适配
	CastType InnerFunc = "cast"

	Lower     InnerFunc = "lower"
	Upper     InnerFunc = "upper"
	Substring InnerFunc = "substring"
	Length    InnerFunc = "length"

	NullIf InnerFunc = "nullif"

	// Concat 可变参数的通用函数
	Concat InnerFunc = "concat"

	Greatest InnerFunc = "greatest"
	Least    InnerFunc = "least"
)

// mysql/doris
const (
	MysqlIfNull InnerFunc = "ifnull"
)

// pg
const (
	CallPgDateTrunc   InnerFunc = "date_trunc"
	CallPGToChar      InnerFunc = "to_char"
	CallPgToDate      InnerFunc = "to_date"
	CallPgToTimestamp InnerFunc = "to_timestamp"

	Cbrt       InnerFunc = "cbrt"
	Trunc      InnerFunc = "trunc"
	PgCoalesce InnerFunc = "coalesce"

	CallPgDateAddDay   InnerFunc = "pg_dateadd_day"
	CallPgDateAddMonth InnerFunc = "pg_dateadd_month"
	CallPgDateAddYear  InnerFunc = "pg_dateadd_year"

	CallPgDateTruncDay   InnerFunc = "pg_datetrunc_day"
	CallPgDateTruncMonth InnerFunc = "pg_datetrunc_month"
	CallPgDateTruncYear  InnerFunc = "pg_datetrunc_year"

	// 由于sql太复杂使用字面量实现
	CallPgDateDiffDay   InnerFunc = "pg_datediff_day"
	CallPgDateDiffMonth InnerFunc = "pg_datediff_month"
	CallPgDateDiffYear  InnerFunc = "pg_datediff_year"
)

// doris
const (
	CallDorisDateTrunc  InnerFunc = "doris_date_trunc"
	CallDorisDateFormat InnerFunc = "date_format"

	CallDorisDateAddDay   InnerFunc = "doris_dateadd_day"
	CallDorisDateAddMonth InnerFunc = "doris_dateadd_month"
	CallDorisDateAddYear  InnerFunc = "doris_dateadd_year"

	CallDorisDateTruncDay   InnerFunc = "doris_datetrunc_day"
	CallDorisDateTruncMonth InnerFunc = "doris_datetrunc_month"
	CallDorisDateTruncYear  InnerFunc = "doris_datetrunc_year"

	CallDorisDateDiffDay   InnerFunc = "doris_datediff_day"
	CallDorisDateDiffMonth InnerFunc = "doris_datediff_month"
	CallDorisDateDiffYear  InnerFunc = "doris_datediff_year"
)

// 通用 format
const (
	// ColumnTypeConvertFormat ColumnTypeConvertAsFormat 转换为指定类型, 占位符1: 列, 占位符2: DDL类型
	ColumnTypeConvertFormat   = `cast(%s AS %s)`
	ColumnTypeConvertAsFormat = `cast(%s AS %s) as "%s"`

	RoundFormat   = `round(%s, %s)`
	RoundFormatAs = `round(%s, %s) as "%s"`

	CeilingFormat   = `ceil(%s)`
	CeilingFormatAs = `ceil(%s) as "%s"`

	FloorFormat   = `floor(%s)`
	FloorFormatAs = `floor(%s) as "%s"`

	AbsFormat       = `abs(%s)`
	AbsAsFormat     = `abs(%s) as "%s"`
	ModFuncFormat   = `mod(%s, %v)`
	ModFuncAsFormat = `mod(%s, %v) as "%s"`
	SqrtFormat      = `sqrt(%s)`
	SqrtAsFormat    = `sqrt(%s) as "%s"`

	LowerFormat       = `lower(%s)`
	LowerAsFormat     = `lower(%s) as "%s"`
	UpperFormat       = `upper(%s)`
	UpperAsFormat     = `upper(%s) as "%s"`
	SubstringFormat   = `substring(%s, %v, %v)`
	SubstringAsFormat = `substring(%s, %v, %v) as "%s"`
	LengthFormat      = `length(%s)`
	LengthAsFormat    = `length(%s) as "%s"`

	ConcatFormat   = `concat(%s)`
	ConcatAsFormat = `concat(%s) as "%s"`

	NullIfFormat = `nullif(%s, %s)`

	GreatestFormat = `greatest(%s)`
	LeastFormat    = `least(%s)`
)

// mysql format
const (
	MysqlIfNullFormat = `ifnull(%s, %s)`
)

// pg format
const (
	// PgDateTruncFormat PgDateTruncAsFormat 转换为离散时间, 最终返回是timestamp类型
	PgDateTruncFormat = `date_trunc(%v, %s)`

	PgDateTruncDayFormat   = `date_trunc('day', %s)`
	PgDateTruncMonthFormat = `date_trunc('month', %s)`
	PgDateTruncYearFormat  = `date_trunc('year', %s)`

	PgDateTruncAsFormat = `date_trunc(%v, %s) as "%s"`

	// ToCharFormat ToCharAsFormat to_char(时间类型字段, format), 最终返回是字符串
	ToCharFormat   = `to_char(%s, %v)`
	ToCharAsFormat = `to_char(%s, %v) as "%s"`

	CallPgToDateFormat        = `to_date(%s, %v)`
	CallPgToDateAsFormat      = `to_date(%s, %v)`
	CallPgToTimestampFormat   = `to_timestamp(%s, %v)`
	CallPgToTimestampAsFormat = `to_timestamp(%s, %v)`

	CbrtFormat    = `cbrt(%s)`
	CbrtAsFormat  = `cbrt(%s) as "%s"`
	TruncFormat   = `trunc(%s, %v)`
	TruncAsFormat = `trunc(%s, %v) as "%s"`

	PgCoalesceFormat = `coalesce(%s)`

	CallPgDateAddDayFormat   = `(%s + INTERVAL '%v DAY')`
	CallPgDateAddMonthFormat = `(%s + INTERVAL '%v MONTH')`
	CallPgDateAddYearFormat  = `(%s + INTERVAL '%v YEAR')`
)

// doris format
const (
	DorisDateTruncFormat = `date_trunc(%s, %v)`

	DorisDateTruncDayFormat   = `date_trunc(%s, 'day')`
	DorisDateTruncMonthFormat = `date_trunc(%s, 'month')`
	DorisDateTruncYearFormat  = `date_trunc(%s, 'year')`

	DorisDateTruncAsFormat = "date_trunc(%s, %v) as `%s`"

	DorisDateFormatFormat = `date_format(%s, %v)`

	DorisDateAddDayFormat   = `DATE_ADD(%s, INTERVAL %v DAY)`
	DorisDateAddMonthFormat = `DATE_ADD(%s, INTERVAL %v MONTH)`
	DorisDateAddYearFormat  = `DATE_ADD(%s, INTERVAL %v YEAR)`

	DorisDateDiffDayFormat   = `DATEDIFF(%s, %s)`
	DorisDateDiffMonthFormat = `TIMESTAMPDIFF(MONTH, %s, %s)`
	DorisDateDiffYearFormat  = `TIMESTAMPDIFF(YEAR, %s, %s)`
)

// InnerFuncFormatMap 内置函数名称和SQL片段 映射
var (
	InnerFuncFormatMap = map[InnerFunc]string{
		// common

		CallRound:   RoundFormat,
		CallCeiling: CeilingFormat,
		CallFloor:   FloorFormat,
		CastType:    ColumnTypeConvertFormat,

		Abs:       AbsFormat,
		ModFunc:   ModFuncFormat,
		Sqrt:      SqrtFormat,
		Lower:     LowerFormat,
		Upper:     UpperFormat,
		Substring: SubstringFormat,
		Length:    LengthFormat,

		Concat:   ConcatFormat,
		Greatest: GreatestFormat,
		Least:    LeastFormat,
		NullIf:   NullIfFormat,

		// 无参数
		NowFunc: NowFormat,

		// mysql

		MysqlIfNull: MysqlIfNullFormat,

		// pg

		CallPgDateTrunc: PgDateTruncFormat,

		CallPgDateTruncDay:   PgDateTruncDayFormat,
		CallPgDateTruncMonth: PgDateTruncMonthFormat,
		CallPgDateTruncYear:  PgDateTruncYearFormat,

		CallPgDateAddDay:   CallPgDateAddDayFormat,
		CallPgDateAddMonth: CallPgDateAddMonthFormat,
		CallPgDateAddYear:  CallPgDateAddYearFormat,
		CallPGToChar:       ToCharFormat,

		CallPgToDate:      CallPgToDateFormat,
		CallPgToTimestamp: CallPgToTimestampFormat,
		Cbrt:              CbrtFormat,
		Trunc:             TruncFormat,

		PgCoalesce: PgCoalesceFormat,

		// doris
		CallDorisDateTrunc: DorisDateTruncFormat,

		CallDorisDateTruncDay:   DorisDateTruncDayFormat,
		CallDorisDateTruncMonth: DorisDateTruncMonthFormat,
		CallDorisDateTruncYear:  DorisDateTruncYearFormat,

		CallDorisDateFormat: DorisDateFormatFormat,

		CallDorisDateAddDay:   DorisDateAddDayFormat,
		CallDorisDateAddMonth: DorisDateAddMonthFormat,
		CallDorisDateAddYear:  DorisDateAddYearFormat,

		CallDorisDateDiffDay:   DorisDateDiffDayFormat,
		CallDorisDateDiffMonth: DorisDateDiffMonthFormat,
		CallDorisDateDiffYear:  DorisDateDiffYearFormat,
	}

	// InnerAsFuncFormatMap 注意: as格式化不再使用, 代码逻辑 as 别名根据驱动做格式化

	InnerAsFuncFormatMap = map[InnerFunc]string{
		// common

		CallRound:   RoundFormatAs,
		CallCeiling: CeilingFormatAs,
		CallFloor:   FloorFormatAs,
		CastType:    ColumnTypeConvertAsFormat,

		Abs:       AbsAsFormat,
		ModFunc:   ModFuncAsFormat,
		Sqrt:      SqrtAsFormat,
		Lower:     LowerAsFormat,
		Upper:     UpperAsFormat,
		Substring: SubstringAsFormat,
		Length:    LengthAsFormat,

		// pg

		CallPgDateTrunc: PgDateTruncAsFormat,
		CallPGToChar:    ToCharAsFormat,

		CallPgToDate:      CallPgToDateAsFormat,
		CallPgToTimestamp: CallPgToTimestampAsFormat,
		Cbrt:              CbrtAsFormat,
		Trunc:             TruncAsFormat,

		// doris
		CallDorisDateTrunc: DorisDateTruncAsFormat,
	}

	// VariadicInnerFuncFormatMap 可变参数
	VariadicInnerFuncFormatMap = map[InnerFunc]string{
		Concat:     ConcatFormat,
		PgCoalesce: PgCoalesceFormat,
		Greatest:   GreatestFormat,
		Least:      LeastFormat,
	}

	NoneArgFuncFormatMap = map[string]string{
		Count1Str: Count1Format,
		NowFunc:   NowFormat,
	}
)

func IsVariadicArgsFunc(exp string) bool {
	return VariadicInnerFuncFormatMap[InnerFunc(exp)] != ""
}

func IsNoneArgFunc(exp string) bool {
	return NoneArgFuncFormatMap[exp] != ""
}
