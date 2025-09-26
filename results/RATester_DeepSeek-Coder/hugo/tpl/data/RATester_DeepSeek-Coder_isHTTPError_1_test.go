package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsHTTPError_1(t *testing.T) {
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				res: &http.Response{
					StatusCode: 200,
				},
			},
			want: false,
		},
		{
			name: "Test 2",
			args: args{
				res: &http.Response{
					StatusCode: 300,
				},
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				res: &http.Response{
					StatusCode: 199,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isHTTPError(tt.args.res); got != tt.want {
				t.Errorf("isHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}
