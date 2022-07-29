package json

import (
	"errors"
	"github.com/json-iterator/go/extra"
	"github.com/yunkeCN/gokit/util"

	"github.com/json-iterator/go"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	bytes, err := Marshal(v)
	return util.Byte2Str(bytes), err
}

func MarshalToStringNoError(v interface{}) string {
	str, _ := MarshalToString(v)
	return str
}

func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(util.Str2Byte(data), v)
}

func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

func UnmarshalString(data string, v interface{}) error {
	return UnmarshalByte(util.Str2Byte(data), v)
}

// UnmarshalStringFuzzyDecoders 解析时容忍空数组作为对象
func UnmarshalStringFuzzyDecoders(data string, v interface{}) error {
	return jsoniter.UnmarshalFromString(data, v)
}
