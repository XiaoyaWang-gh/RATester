package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := plainBinding{}
	obj := make(map[string]interface{})
	body := []byte(`{"name":"test"}`)
	err := b.BindBody(body, &obj)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if obj["name"] != "test" {
		t.Errorf("Expected 'test', got %v", obj["name"])
	}
}
