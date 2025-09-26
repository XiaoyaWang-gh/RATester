package request

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	content := []byte("test content")
	r.Body(content)
	if !reflect.DeepEqual(r.body, content) {
		t.Errorf("Expected %v, got %v", content, r.body)
	}
}
