package common

import (
	"github.com/forhsd/logger"
	"github.com/ouseikou/sqlbuilder/common/clause"
	"github.com/ouseikou/sqlbuilder/common/facade"
	pb "github.com/ouseikou/sqlbuilder/gen/proto"
)

// todo 1. java Object传参顺序问题影响SQL生成和结果 2. schema.table.column 参考MB 3.format在占位符写死了,根据方言决定 "" 或``

func BuildModelSqlByJson(request clause.BuilderRequest) (string, error) {
	builder, bErr := facade.RunBuilder(request)

	defer func() {
		if err := recover(); bErr != nil {
			logger.Error("构建SQL失败, 原因: %v", err)
		}
	}()

	boundSQL, sqlErr := builder.ToBoundSQL()
	if sqlErr != nil {
		logger.Error("转换SQL失败, 原因: %v", sqlErr)
	}

	logger.Debug("执行模型模式SQL构建策略, pg方言SQL: \n %s", boundSQL)

	return boundSQL, nil
}

func BuildTemplateSqlByJson(request clause.BuilderRequest) (string, error) {
	logger.Debug("执行模板模式SQL构建策略, pg方言SQL: %s", "SELECT 模板")
	return "", nil
}

// BuildSqlByModelProto 使用 proto 解决 http 传输联合类型接收为 map 的元组类型擦除问题
func BuildSqlByModelProto(request *pb.BuilderRequest) (string, error) {
	// 根据方言创建facade
	f, err := facade.CreateModelBuilderFacade(request.Driver)
	if err != nil {
		return "", err
	}
	// 构建 xormBuilder
	xormBuilder, err := f.BuildSQL(request)
	if err != nil {
		return "", err
	}
	// builder -> sql
	boundSQL, err := xormBuilder.ToBoundSQL()
	if err != nil {
		return "", err
	}

	logger.Debug("执行模型模式SQL构建策略, SQL: \n %s", boundSQL)

	return boundSQL, nil
}

func BuildSqlByTemplateProto(request *pb.BuilderRequest) (string, error) {
	logger.Debug("执行模板模式SQL构建策略, pg方言SQL: %s", "SELECT 模板")
	return "", nil
}
