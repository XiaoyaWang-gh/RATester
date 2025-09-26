package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlGet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	ctrl.CtrlGet("/get", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
		ctx.Output.Body([]byte("test"))
	})

	req, _ := http.NewRequest("GET", "/test/get", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	if w.Body.String() != "test" {
		t.Errorf("Expected body 'test', got %s", w.Body.String())
	}
}
