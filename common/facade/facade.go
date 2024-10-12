package facade

import (
	"errors"

	pb "github.com/ouseikou/sqlbuilder/gen/proto"
	xorm "xorm.io/builder"
)

// -----------------------------------------------   facade: 通过模型构建sql  --------------------------------------------

// CreateModelBuilderFacade 工厂方法根据Driver选择具体Facade
// 参数:
//   - driver: 驱动
//
// 返回值:
//   - 模型模式SqlBuilder门面接口
//   - 异常
func CreateModelBuilderFacade(driver pb.Driver) (ModelBuilderFacade, error) {
	switch driver {
	case pb.Driver_DRIVER_POSTGRES:
		return &PostgresModelBuilderFacade{}, nil
	default:
		return nil, errors.New("driver not support")
	}
}

// ModelBuilderFacade 定义Facade接口(模型模式SqlBuilder门面接口)
// 参数:
//   - request: proto 请求对象
//
// 返回值:
//   - 返回 sql
//   - 异常
type ModelBuilderFacade interface {
	BuildSQL(request *pb.BuilderRequest) (string, error)

	//BuildDialect() (*xorm.Builder, error)
	//BuildBasic(dialectBuilder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder
	//BuildOther(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder
}

// AbstractModelBuilderFacade 定义抽象Facade结构体，提供模板方法模式的基础实现
// 模板方法步骤:
// 1. BuildDialect, 创建方言
// 2. BuildBasic, 链式创建 select & from
// 3. BuildOther, 链式创建 SQL 其他关键字
type AbstractModelBuilderFacade struct {
	//act ModelBuilderFacade
}

// BuildSQL 抽象实现(模板方法模式)
func (f *AbstractModelBuilderFacade) BuildSQL(request *pb.BuilderRequest) (string, error) {
	impl, err := CreateModelBuilderFacade(request.Driver)
	if err != nil {
		return "", err
	}

	return impl.BuildSQL(request)
}

// BuildDialect 抽象方法: 构建方言, 具体实现由子类提供
// 参数:
//
// 返回值:
//   - 返回一个 xorm.Builder 对象
//   - 异常
func (f *AbstractModelBuilderFacade) BuildDialect() (*xorm.Builder, error) {
	return &xorm.Builder{}, nil
}

// BuildBasic 抽象方法: 构建 select & from, 具体实现由子类提供
// 参数:
//   - appendBuilder: 上一次的 xorm.Builder 指针
//   - dialectBuilder: xorm.Builder 对象
//   - sqlRef: 当前层级的sqlRef对象
//
// 返回值:
//   - 返回一个 xorm.Builder 对象
func (f *AbstractModelBuilderFacade) BuildBasic(appendBuilder *xorm.Builder, dialectBuilder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	return &xorm.Builder{}
}

// BuildOther 抽象方法: 构建SQL其他关键字内容，具体实现由子类提供
// 参数:
//   - builder: xorm.Builder 指针对象
//   - sqlRef: 当前层级的sqlRef对象
//
// 返回值:
//   - 返回一个 xorm.Builder 对象
func (f *AbstractModelBuilderFacade) BuildOther(builder *xorm.Builder, sqlRef *pb.SqlReference) *xorm.Builder {
	return &xorm.Builder{}
}

// -----------------------------------------------   facade: 通过模板构建sql  --------------------------------------------

// BuildSqlByTemplate : 根据模板语法构建sql
// 参数:
//   - request: sqlbuilder 请求体
//
// 返回值:
//   - 返回 sql
//   - 返回 错误信息
func BuildSqlByTemplate(request *pb.BuilderRequest) (string, error) {
	return RenderMasterTemplate(request.GetBuilders()[0].GetSql().GetTemplate())
}
