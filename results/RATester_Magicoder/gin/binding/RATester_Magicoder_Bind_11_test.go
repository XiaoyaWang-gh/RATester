package binding

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Field1 string `yaml:"field1"`
		Field2 int    `yaml:"field2"`
	}

	testObj := testStruct{}
	req, _ := http.NewRequest("POST", "/", strings.NewReader("field1: test\nfield2: 123\n"))

	err := yamlBinding{}.Bind(req, &testObj)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if testObj.Field1 != "test" {
		t.Errorf("Expected Field1 to be 'test', but got %v", testObj.Field1)
	}

	if testObj.Field2 != 123 {
		t.Errorf("Expected Field2 to be 123, but got %v", testObj.Field2)
	}
}
