package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestPut_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	handlerCalled := false
	handler := func(ctx *beecontext.Context) {
		handlerCalled = true
	}

	app.Put("/test", handler)

	req, err := http.NewRequest("PUT", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	if !handlerCalled {
		t.Errorf("handler was not called")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
