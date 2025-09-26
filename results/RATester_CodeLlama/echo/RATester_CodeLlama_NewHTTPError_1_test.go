package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPError_1(t *testing.T) {
	type args struct {
		code    int
		message []interface{}
	}
	tests := []struct {
		name string
		args args
		want *HTTPError
	}{
		{
			name: "test case 1",
			args: args{
				code:    404,
				message: []interface{}{"test message"},
			},
			want: &HTTPError{
				Code:    404,
				Message: "test message",
			},
		},
		{
			name: "test case 2",
			args: args{
				code:    404,
				message: []interface{}{"test message"},
			},
			want: &HTTPError{
				Code:    404,
				Message: "test message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewHTTPError(tt.args.code, tt.args.message...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}
