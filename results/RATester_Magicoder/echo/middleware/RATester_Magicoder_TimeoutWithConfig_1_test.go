package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestTimeoutWithConfig_1(t *testing.T) {
	type args struct {
		config TimeoutConfig
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

			if got := TimeoutWithConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeoutWithConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
