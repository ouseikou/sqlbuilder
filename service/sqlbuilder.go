package service

import (
	"errors"

	"github.com/ouseikou/sqlbuilder/common"
	"github.com/ouseikou/sqlbuilder/common/clause"
	pb "github.com/ouseikou/sqlbuilder/gen/proto"
)

// BuildSqlByJson http协议
func BuildSqlByJson(builder clause.BuilderRequest) (string, error) {
	mode := builder.Strategy
	// service 根据构建策略决定具体实现, sql生成根据driver方言来生成
	switch mode {
	case clause.ModelStrategy:
		sql, err := common.BuildModelSqlByJson(builder)
		return sql, err
	case clause.TemplateStrategy:
		sql, err := common.BuildTemplateSqlByJson(builder)
		return sql, err
	default:
		return "", errors.New("未知SQL构建策略")
	}
}

// BuildSqlByProto proto 协议
func BuildSqlByProto(builder *pb.BuilderRequest) (string, error) {
	mode := builder.Strategy
	// service 根据构建策略决定具体实现, sql生成根据driver方言来生成
	switch mode {
	case pb.BuilderStrategy_STRATEGY_MODEL:
		sql, err := common.BuildSqlByModelProto(builder)
		return sql, err
	case pb.BuilderStrategy_STRATEGY_TEMPLATE:
		sql, err := common.BuildSqlByTemplateProto(builder)
		return sql, err
	default:
		return "", errors.New("未知SQL构建策略")
	}
}
