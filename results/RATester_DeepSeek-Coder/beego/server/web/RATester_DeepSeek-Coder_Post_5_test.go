package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestPost_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	testHandler := func(ctx *beecontext.Context) {
		ctx.Output.SetStatus(200)
		ctx.Output.Body([]byte("test"))
	}

	n.Post("/post", testHandler)

	req, err := http.NewRequest("POST", "/test/post", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "test"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
