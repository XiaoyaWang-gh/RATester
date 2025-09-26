package binding

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := jsonBinding{}

	// Test case 1: Valid request
	req, _ := http.NewRequest("POST", "/", strings.NewReader(`{"name":"John","age":30}`))
	obj := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	err := b.Bind(req, &obj)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if obj.Name != "John" || obj.Age != 30 {
		t.Errorf("Expected obj to be %v, but got %v", obj, `{"name":"John","age":30}`)
	}

	// Test case 2: Invalid request
	req, _ = http.NewRequest("POST", "/", nil)
	err = b.Bind(req, &obj)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 3: Invalid JSON
	req, _ = http.NewRequest("POST", "/", strings.NewReader(`{"name":"John","age":"thirty"}`))
	err = b.Bind(req, &obj)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
