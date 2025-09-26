package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestHead_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	testFunc := func(ctx *beecontext.Context) {
		ctx.Output.SetStatus(200)
	}

	app.Head("/test", testFunc)

	req, err := http.NewRequest("HEAD", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.DefaultServeMux.HandleFunc("/", app.Handlers.ServeHTTP)
	http.DefaultServeMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
