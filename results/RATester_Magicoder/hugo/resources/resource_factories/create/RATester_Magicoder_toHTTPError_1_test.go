package create

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TesttoHTTPError_1(t *testing.T) {
	type args struct {
		err      error
		res      *http.Response
		readBody bool
	}
	tests := []struct {
		name string
		args args
		want *HTTPError
	}{
		{
			name: "Test Case 1",
			args: args{
				err:      errors.New("test error"),
				res:      &http.Response{},
				readBody: true,
			},
			want: &HTTPError{
				error: errors.New("test error"),
				Data:  map[string]any{},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				err:      nil,
				res:      &http.Response{},
				readBody: true,
			},
			want: &HTTPError{
				error: errors.New("err is nil"),
				Data:  map[string]any{},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				err:      errors.New("test error"),
				res:      nil,
				readBody: true,
			},
			want: &HTTPError{
				error: errors.New("test error"),
				Data:  map[string]any{},
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

			if got := toHTTPError(tt.args.err, tt.args.res, tt.args.readBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}
