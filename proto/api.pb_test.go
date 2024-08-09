package proto

import (
	"fmt"
	"testing"

	"github.com/ouseikou/sqlbuilder/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

var protoJsonOP = protojson.MarshalOptions{
	EmitUnpopulated: true,
	// UseEnumNumbers:  true,
}

func makeSelect1Vars() []*proto.MixVars {
	vars := make([]*proto.MixVars, 0)
	col := &proto.MixVars_Column{
		Column: &proto.Column{Schema: "sample_data", Table: "products", Field: "price", Alias: "合计", AggAble: false, UseAs: false},
	}
	vars = append(vars, &proto.MixVars{Vars: col})

	return vars
}

func makeSelect2Vars() []*proto.MixVars {
	vars := make([]*proto.MixVars, 0)
	col := &proto.MixVars_Column{
		Column: &proto.Column{Schema: "sample_data", Table: "products", Field: "created_at", Alias: "created_at表达式内部as", AggAble: false, UseAs: false},
	}
	colStr := &proto.MixVars_Context{
		Context: "month",
	}
	vars = append(vars, &proto.MixVars{Vars: colStr}, &proto.MixVars{Vars: col})

	return vars
}

func makeSelect() []*proto.MixField {
	selects := make([]*proto.MixField, 0)
	selectCol1 := &proto.MixField_Column{
		Column: &proto.Column{Schema: "sample_data", Table: "products", Field: "category", Alias: "分类", AggAble: false, UseAs: true},
	}
	selectExpr1 := &proto.MixField_Expression{
		Expression: &proto.Expression{
			Call:     "sum",
			CallType: proto.CallType_CALL_TYPE_AGG,
			CallAs:   "合计价格",
			UseAs:    true,
			Vars:     makeSelect1Vars(),
		},
	}

	selectExpr2 := &proto.MixField_Expression{
		Expression: &proto.Expression{
			Call:     "DATE_TRUNC",
			CallType: proto.CallType_CALL_TYPE_INNER,
			CallAs:   "离散日期(月)",
			UseAs:    true,
			Vars:     makeSelect2Vars(),
		},
	}

	selects = append(selects, &proto.MixField{Mix: selectCol1}, &proto.MixField{Mix: selectExpr1}, &proto.MixField{Mix: selectExpr2})

	return selects
}

func makeGroupByVars() []*proto.MixVars {
	vars := make([]*proto.MixVars, 0)
	col := &proto.MixVars_Column{
		Column: &proto.Column{Schema: "sample_data", Table: "products", Field: "created_at", Alias: "created_at表达式内部as", AggAble: false, UseAs: false},
	}
	colStr := &proto.MixVars_Context{
		Context: "month",
	}
	vars = append(vars, &proto.MixVars{Vars: colStr}, &proto.MixVars{Vars: col})

	return vars
}

func makeGroupBy() []*proto.MixField {
	selects := make([]*proto.MixField, 0)
	groupByCol1 := &proto.MixField_Column{
		Column: &proto.Column{Schema: "sample_data", Table: "products", Field: "category", Alias: "分类", AggAble: false, UseAs: false},
	}

	groupByExpr2 := &proto.MixField_Expression{
		Expression: &proto.Expression{
			Call:     "DATE_TRUNC",
			CallType: proto.CallType_CALL_TYPE_INNER,
			CallAs:   "离散日期(月)",
			UseAs:    false,
			Vars:     makeGroupByVars(),
		},
	}

	selects = append(selects, &proto.MixField{Mix: groupByExpr2}, &proto.MixField{Mix: groupByCol1})

	return selects
}

func TestBuilderRequest_Descriptor(t *testing.T) {
	sql := &proto.SqlReference{
		From:        &proto.Table{TableName: "products", TableSchema: "sample_data", TableAlias: "products"},
		Join:        []*proto.Join{},
		Where:       []*proto.Expression{},
		GroupBy:     makeGroupBy(),
		Aggregation: []*proto.Expression{},
		Select:      makeSelect(),
		OrderBy:     []*proto.OrderBy{},
		Limit:       &proto.Limit{},
	}

	modelSql := &proto.MixSql_Model{Model: sql}

	deep0 := &proto.DeepWrapper{Deep: 0, Sql: &proto.MixSql{Ref: modelSql}}

	request := &proto.BuilderRequest{
		Driver:   proto.Driver_DRIVER_POSTGRES,
		Strategy: proto.BuilderStrategy_BUILDER_STRATEGY_MODEL,
		Builders: []*proto.DeepWrapper{deep0},
	}

	buf, err := protoJsonOP.Marshal(request)

	if err != nil {
		fmt.Println(err)
	}

	json := string(buf)
	fmt.Printf(json)

	t.Run("测试", func(t *testing.T) {
		if len(request.Builders[0].Sql.GetModel().GroupBy) != 2 && len(request.Builders[0].Sql.GetModel().Select) != 3 {
			t.Errorf("断言失败")
		}
	})
}
