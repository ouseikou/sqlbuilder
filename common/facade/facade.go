package facade

import (
	"errors"

	pb "github.com/ouseikou/sqlbuilder/gen/proto"
	xorm "xorm.io/builder"
)

// CreateModelBuilderFacade 工厂方法根据Driver选择具体Facade
func CreateModelBuilderFacade(driver pb.Driver) (ModelBuilderFacade, error) {
	switch driver {
	case pb.Driver_DRIVER_POSTGRES:
		return &PostgresModelBuilderFacade{}, nil
	default:
		return nil, errors.New("driver not support")
	}
}

// ModelBuilderFacade 定义Facade接口
type ModelBuilderFacade interface {
	BuildSQL(request *pb.BuilderRequest) (*xorm.Builder, error)

	//BuildDialect(request *pb.BuilderRequest) (*xorm.Builder, error)
	//BuildBasic(dialectBuilder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder
	//BuildOther(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder
}

// AbstractModelBuilderFacade 定义抽象Facade结构体，提供模板方法模式的基础实现
type AbstractModelBuilderFacade struct {
	//act ModelBuilderFacade
}

// BuildSQL 抽象实现(模板方法模式)
func (f *AbstractModelBuilderFacade) BuildSQL(request *pb.BuilderRequest) (*xorm.Builder, error) {
	impl, err := CreateModelBuilderFacade(request.Driver)
	if err != nil {
		return nil, err
	}

	return impl.BuildSQL(request)
}

// BuildDialect 抽象方法，具体实现由子类提供
func (f *AbstractModelBuilderFacade) BuildDialect(request *pb.BuilderRequest) (*xorm.Builder, error) {
	driver := request.Driver
	switch driver {
	case pb.Driver_DRIVER_POSTGRES:
		return xorm.Dialect(xorm.POSTGRES), nil
	default:
		return nil, errors.New("未知SQL构建策略")
	}
}

func (f *AbstractModelBuilderFacade) BuildBasic(dialectBuilder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	return &xorm.Builder{}
}

func (f *AbstractModelBuilderFacade) BuildOther(builder *xorm.Builder, sqlRef *pb.SQLReference) *xorm.Builder {
	return &xorm.Builder{}
}
