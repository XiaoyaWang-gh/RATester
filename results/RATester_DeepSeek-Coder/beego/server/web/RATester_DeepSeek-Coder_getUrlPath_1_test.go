package web

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestGetUrlPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		cfg: &Config{
			RouterCaseSensitive: true,
		},
	}

	ctx := &beecontext.Context{
		Request: &http.Request{
			URL: &url.URL{
				Path: "/Hello",
			},
		},
	}

	expected := "/Hello"
	result := ctrl.getUrlPath(ctx)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
