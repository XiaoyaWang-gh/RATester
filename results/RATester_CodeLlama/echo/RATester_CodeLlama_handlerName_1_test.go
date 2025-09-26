package echo

import (
	"fmt"
	"testing"
)

func TestHandlerName_1(t *testing.T) {
	type args struct {
		h HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				h: func(c Context) error {
					return nil
				},
			},
			want: "github.com/labstack/echo/v4.(*Echo).ServeHTTP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := handlerName(tt.args.h); got != tt.want {
				t.Errorf("handlerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
