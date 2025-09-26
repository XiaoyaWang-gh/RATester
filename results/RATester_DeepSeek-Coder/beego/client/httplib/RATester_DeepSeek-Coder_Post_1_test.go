package httplib

import (
	"fmt"
	"testing"
)

func TestPost_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *BeegoHTTPRequest
	}{
		{
			name: "Test Post",
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

			got := Post(tt.args.url)
			if got.url != tt.want.url {
				t.Errorf("Post() = %v, want %v", got, tt.want)
			}
		})
	}
}
