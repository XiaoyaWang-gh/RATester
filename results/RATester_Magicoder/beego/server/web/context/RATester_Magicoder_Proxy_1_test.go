package context

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"X-Forwarded-For": []string{"192.168.1.1, 192.168.1.2"},
				},
			},
		},
	}

	expected := []string{"192.168.1.1", "192.168.1.2"}
	result := input.Proxy()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
