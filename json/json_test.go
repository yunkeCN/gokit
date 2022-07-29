package json

import (
	"fmt"
	"testing"
)

var testMap = map[string]interface{}{
	"name": "kael",
	"age":  18,
	"friend": map[string]interface{}{
		"name": "kael2",
		"age":  8,
	},
}

type TestPerson struct {
	Name   string
	Age    int
	Friend struct {
		Name string
		Age  int
	}
}

func TestMarshal(t *testing.T) {
	fmt.Println(Marshal(testMap))
}

func TestMarshalToString(t *testing.T) {

	var mm = map[string]int64{
		"a": 111,
		"s": 222222222222222,
	}

	fmt.Println(MarshalToString(mm))
}

func TestMarshalToStringNoError(t *testing.T) {
	fmt.Println(MarshalToStringNoError(testMap))
}

func TestUnmarshal(t *testing.T) {
	var tp = &TestPerson{}
	Unmarshal("{\"age\":18,\"friend\":{\"age\":8,\"name\":\"kael2\"},"+
		"\"name\":\"kael\"}", tp)
	fmt.Println(tp)
}

func TestUnmarshalStringFuzzyDecoders(t *testing.T) {
	var temp = `{"a":[],"b":"hello"}`

	var s struct {
		A []int  `json:"a"`
		B string `json:"b"`
	}
	err := UnmarshalStringFuzzyDecoders(temp, &s)
	if err != nil {
		t.Error(err)
	}

	//result is {[], hello}
	t.Logf("%v", s)
}
