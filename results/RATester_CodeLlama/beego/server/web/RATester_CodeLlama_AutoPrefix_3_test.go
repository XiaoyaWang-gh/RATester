package web

import (
	"fmt"
	"testing"
)

func TestAutoPrefix_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{}
	app.AutoPrefix("/v1", &Controller{})
}
