package common

type BuilderStrategy int
type Driver int

const (
	// ModelStrategy 模型模式
	ModelStrategy BuilderStrategy = iota
	// TemplateStrategy 模板模式
	TemplateStrategy
)

const (
	DriverPostgres Driver = iota
	DriverMysql
	DriverDoris
)

func (i BuilderStrategy) EnType() string {
	names := [...]string{"model", "template"}
	if i < 0 || int(i) >= len(names) {
		// 处理未知类型的情况
		return "unknown"
	}
	return names[i]
}

func BuilderStrategyEnumType(index BuilderStrategy) string {
	switch index {
	case ModelStrategy, TemplateStrategy:
		return index.EnType()
	default:
		return ModelStrategy.EnType()
	}
}

func (i Driver) EnType() string {
	names := [...]string{"postgres", "mysql", "doris"}
	if i < 0 || int(i) >= len(names) {
		// 处理未知类型的情况
		return "unknown"
	}
	return names[i]
}

func DriverEnumType(index Driver) string {
	switch index {
	case DriverPostgres, DriverMysql, DriverDoris:
		return index.EnType()
	default:
		return "unknown"
	}
}
