package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlPut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.CtrlPut("/test", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	})

	req := httptest.NewRequest("PUT", "/test", nil)
	w := httptest.NewRecorder()
	ctrl.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, res.StatusCode)
	}
}
