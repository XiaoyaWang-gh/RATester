package web

import (
	"fmt"
	"testing"
)

func TestOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f HandleFunc
	var app *HttpServer
	app = &HttpServer{}
	app.Options(rootpath, f)
}
