package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestHead_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	pattern := "/test"
	handler := func(ctx *beecontext.Context) {
		ctx.Output.SetStatus(200)
	}
	ctrl.Head(pattern, handler)

	req, err := http.NewRequest("HEAD", pattern, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ctrl.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
