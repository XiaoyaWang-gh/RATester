package web

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestInput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{
			Request: &http.Request{
				Form: url.Values{
					"key": []string{"value"},
				},
			},
		},
	}

	values, err := ctrl.Input()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if values.Get("key") != "value" {
		t.Errorf("Expected 'value', got %v", values.Get("key"))
	}
}
