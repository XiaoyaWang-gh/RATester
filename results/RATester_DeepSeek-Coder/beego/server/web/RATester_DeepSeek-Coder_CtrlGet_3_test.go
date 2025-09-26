package web

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlGet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.CtrlGet("/test", func(ctx *context.Context) {
		ctx.Output.Body([]byte("test"))
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	ctrl.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	if string(data) != "test" {
		t.Errorf("expected body to be 'test', got %s", string(data))
	}
}
