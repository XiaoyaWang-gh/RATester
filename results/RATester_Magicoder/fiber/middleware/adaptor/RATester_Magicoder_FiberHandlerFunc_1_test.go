package adaptor

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestFiberHandlerFunc_1(t *testing.T) {
	type args struct {
		h fiber.Handler
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
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

			if got := FiberHandlerFunc(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FiberHandlerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
