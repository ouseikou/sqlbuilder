package service

import (
	"errors"
	"sql-builder/common"
	"sql-builder/common/clause"
	"sql-builder/common/strategy"
)

func BuildSql(builder clause.BuilderRequest) (string, error) {
	mode := builder.Strategy
	// service 根据构建策略决定具体实现, sql生成根据driver方言来生成
	switch mode {
	case common.ModelStrategy:
		sql, err := strategy.BuilderModelSql(builder)
		return sql, err
	case common.TemplateStrategy:
		sql, err := strategy.BuilderTemplateSql(builder)
		return sql, err
	default:
		return "", errors.New("未知SQL构建策略")
	}
}
