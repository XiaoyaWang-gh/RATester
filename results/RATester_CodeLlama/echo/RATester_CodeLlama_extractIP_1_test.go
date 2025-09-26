package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExtractIP_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				req: &http.Request{
					RemoteAddr: "127.0.0.1:8080",
				},
			},
			want: "127.0.0.1",
		},
		{
			name: "case 2",
			args: args{
				req: &http.Request{
					RemoteAddr: "[::1]:8080",
				},
			},
			want: "::1",
		},
		{
			name: "case 3",
			args: args{
				req: &http.Request{
					RemoteAddr: "127.0.0.1",
				},
			},
			want: "127.0.0.1",
		},
		{
			name: "case 4",
			args: args{
				req: &http.Request{
					RemoteAddr: "[::1]",
				},
			},
			want: "::1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := extractIP(tt.args.req); got != tt.want {
				t.Errorf("extractIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
