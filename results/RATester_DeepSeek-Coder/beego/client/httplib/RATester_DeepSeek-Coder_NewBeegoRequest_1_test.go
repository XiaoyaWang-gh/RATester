package httplib

import (
	"fmt"
	"testing"
)

func TestNewBeegoRequest_1(t *testing.T) {
	type args struct {
		rawurl string
		method string
	}
	tests := []struct {
		name string
		args args
		want *BeegoHTTPRequest
	}{
		{
			name: "TestNewBeegoRequest",
			args: args{
				rawurl: "http://example.com",
				method: "GET",
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

			got := NewBeegoRequest(tt.args.rawurl, tt.args.method)
			if got.url != tt.want.url {
				t.Errorf("NewBeegoRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
