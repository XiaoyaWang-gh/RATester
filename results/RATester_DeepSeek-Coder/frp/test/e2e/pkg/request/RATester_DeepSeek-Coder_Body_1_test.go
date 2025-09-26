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
	expected := &Request{body: content}

	result := r.Body(content)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
