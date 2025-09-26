package middleware

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

func TestContextTimeout_1(t *testing.T) {
	type args struct {
		timeout time.Duration
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

			if got := ContextTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContextTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
