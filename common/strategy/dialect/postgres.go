package dialect

import (
	"github.com/forhsd/logger"
	"github.com/ouseikou/sqlbuilder/common/clause"
	"github.com/ouseikou/sqlbuilder/util"
)

func BuildModelSql(request clause.BuilderRequest) (string, error) {

	// todo 1. grpc转HTTP处理元组类型擦除 2. java Object传参顺序问题影响SQL生成和结果 3. schema.table.column 参考MB 4.format在占位符写死了,根据方言决定 "" 或``
	builder, bErr := util.RunBuilder(request)

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

func BuildTemplateSql(request clause.BuilderRequest) (string, error) {

	logger.Debug("执行模板模式SQL构建策略, pg方言SQL: %s", "SELECT 模板")
	return "", nil
}
