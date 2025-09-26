package binding

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	binding := formPostBinding{}

	type TestStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testStruct := TestStruct{}

	req, err := http.NewRequest("POST", "/test", strings.NewReader("field1=test&field2=123"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = binding.Bind(req, &testStruct)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if testStruct.Field1 != "test" {
		t.Errorf("Expected Field1 to be 'test', but got %s", testStruct.Field1)
	}

	if testStruct.Field2 != 123 {
		t.Errorf("Expected Field2 to be 123, but got %d", testStruct.Field2)
	}
}
