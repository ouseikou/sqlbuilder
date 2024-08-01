package util

import "encoding/json"

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

func ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
