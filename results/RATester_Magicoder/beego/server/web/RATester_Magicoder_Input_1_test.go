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

	form, err := ctrl.Input()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if form == nil {
		t.Error("Expected form to not be nil")
	}

	if len(form) != 1 {
		t.Errorf("Expected form to have 1 element, but got %d", len(form))
	}

	if form.Get("key") != "value" {
		t.Errorf("Expected form to have key 'value', but got %s", form.Get("key"))
	}
}
