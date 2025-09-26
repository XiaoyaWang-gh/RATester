package web

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestGetFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{
			Request: &http.Request{
				Form: make(url.Values),
			},
		},
	}

	ctrl.Ctx.Request.Form.Add("key", "value")

	file, header, err := ctrl.GetFile("key")
	if err != nil {
		t.Errorf("GetFile returned an error: %v", err)
	}

	if file == nil {
		t.Error("GetFile returned nil file")
	}

	if header == nil {
		t.Error("GetFile returned nil header")
	}
}
