package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestrequestHeader_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *Context
		args args
		want string
	}{
		{
			name: "Test Case 1",
			c: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
				},
			},
			args: args{key: "Content-Type"},
			want: "application/json",
		},
		{
			name: "Test Case 2",
			c: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/xml"},
					},
				},
			},
			args: args{key: "Content-Type"},
			want: "application/xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.requestHeader(tt.args.key); got != tt.want {
				t.Errorf("requestHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
