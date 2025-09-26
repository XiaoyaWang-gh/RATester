package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlHead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	ctrl.CtrlHead("/head", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	})

	req, _ := http.NewRequest("HEAD", "/test/head", nil)
	w := httptest.NewRecorder()

	ctrl.handlers.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
