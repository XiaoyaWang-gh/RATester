package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your handler logic here
	})

	app.Handler("/test", handler)

	// Add your assertions here
}
