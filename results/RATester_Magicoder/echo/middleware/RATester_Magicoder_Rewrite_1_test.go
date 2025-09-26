package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRewrite_1(t *testing.T) {
	type args struct {
		rules map[string]string
	}
	tests := []struct {
		name string
		args args
		want echo.MiddlewareFunc
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

			if got := Rewrite(tt.args.rules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rewrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
