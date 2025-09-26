package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApplyMiddleware_1(t *testing.T) {
	type args struct {
		h          HandlerFunc
		middleware []MiddlewareFunc
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
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

			if got := applyMiddleware(tt.args.h, tt.args.middleware...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
