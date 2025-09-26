package httplib

import (
	"fmt"
	"testing"
)

func TestGet_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *BeegoHTTPRequest
	}{
		{
			name: "TestGet",
			args: args{
				url: "http://example.com",
			},
			want: &BeegoHTTPRequest{
				url: "http://example.com",
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

			got := Get(tt.args.url)
			if got.url != tt.want.url {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
