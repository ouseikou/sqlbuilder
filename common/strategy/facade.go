package strategy

import (
	"errors"

	"github.com/ouseikou/sqlbuilder/common"
	"github.com/ouseikou/sqlbuilder/common/clause"
	"github.com/ouseikou/sqlbuilder/common/strategy/dialect"
)

func BuilderModelSql(builder clause.BuilderRequest) (string, error) {
	driver := builder.Driver
	// facade, 根据driver的方言决定如何生成SQL
	switch driver {
	case common.DriverPostgres:
		sql, _ := dialect.BuildModelSql(builder)
		return sql, nil
	default:
		return "", errors.New("未知SQL构建策略")
	}
}

func BuilderTemplateSql(builder clause.BuilderRequest) (string, error) {
	driver := builder.Driver
	// facade, 根据driver的方言决定如何生成SQL
	switch driver {
	case common.DriverPostgres:
		sql, _ := dialect.BuildTemplateSql(builder)
		return sql, nil
	default:
		return "", errors.New("未知SQL构建策略")
	}
}
