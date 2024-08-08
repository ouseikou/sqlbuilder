package util

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
)

// ProtoJsonOperate proto 序列化保留零值和枚举零值
var ProtoJsonOperate = protojson.MarshalOptions{
	EmitUnpopulated: true,
	// UseEnumNumbers:  true,
}

func Serialize(data interface{}) ([]byte, error) {
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}

func Deserialize(serialized []byte, target interface{}) error {
	err := json.Unmarshal(serialized, target)
	if err != nil {
		return err
	}
	return nil
}

func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
