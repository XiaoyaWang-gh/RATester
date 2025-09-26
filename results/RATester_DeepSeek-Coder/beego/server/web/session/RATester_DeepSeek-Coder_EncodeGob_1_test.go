package session

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeGob_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string
		Age  int
	}
	testMap := make(map[interface{}]interface{})
	testMap["testKey"] = testStruct{"John", 30}
	expected, _ := json.Marshal(testMap)
	result, err := EncodeGob(testMap)
	if err != nil {
		t.Errorf("EncodeGob() error = %v, wantErr %v", err, nil)
		return
	}
	actual, _ := json.Marshal(result)
	if string(actual) != string(expected) {
		t.Errorf("EncodeGob() = %v, want %v", actual, expected)
	}
}
