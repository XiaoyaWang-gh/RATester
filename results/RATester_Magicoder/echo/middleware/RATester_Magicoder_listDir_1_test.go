package middleware

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestListDir_1(t *testing.T) {
	type args struct {
		t    *template.Template
		name string
		dir  http.File
		res  *echo.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := listDir(tt.args.t, tt.args.name, tt.args.dir, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("listDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
