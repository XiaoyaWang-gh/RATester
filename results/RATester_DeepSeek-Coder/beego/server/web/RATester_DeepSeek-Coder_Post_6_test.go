package web

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestPost_6(t *testing.T) {
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
		ctx.Output.Body([]byte("test"))
	}

	app.Post("/test", testFunc)

	req := httptest.NewRequest("POST", "/test", nil)
	w := httptest.NewRecorder()

	app.Server.Handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if string(data) != "test" {
		t.Errorf("Expected 'test', got '%s'", string(data))
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", res.StatusCode)
	}
}
