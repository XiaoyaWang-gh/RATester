package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBindForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Request: &http.Request{
			Form: url.Values{
				"key": []string{"value"},
			},
		},
	}

	obj := struct {
		Key string
	}{}

	err := ctx.BindForm(&obj)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if obj.Key != "value" {
		t.Errorf("Expected obj.Key to be 'value', but got %v", obj.Key)
	}
}
