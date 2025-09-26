package httplib

import (
	"fmt"
	"testing"
)

func TestYAMLBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testObj := testStruct{
		Name: "John Doe",
		Age:  30,
	}

	b := &BeegoHTTPRequest{}
	_, err := b.YAMLBody(&testObj)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedBody := `name: John Doe
age: 30
`
	if string(b.body) != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, string(b.body))
	}

	if b.req.Header.Get("Content-Type") != "application/x+yaml" {
		t.Errorf("Expected Content-Type to be 'application/x+yaml', got %v", b.req.Header.Get("Content-Type"))
	}
}
