package helpers

import (
	"fmt"
	"testing"
)

func TestPermalinkForBaseURL_1(t *testing.T) {
	type args struct {
		link    string
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				link:    "/test",
				baseURL: "https://example.com",
			},
			want: "https://example.com/test",
		},
		{
			name: "Test case 2",
			args: args{
				link:    "/test/sub",
				baseURL: "https://example.com/",
			},
			want: "https://example.com/test/sub",
		},
		{
			name: "Test case 3",
			args: args{
				link:    "/test/sub/",
				baseURL: "https://example.com/",
			},
			want: "https://example.com/test/sub/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &PathSpec{}
			if got := p.PermalinkForBaseURL(tt.args.link, tt.args.baseURL); got != tt.want {
				t.Errorf("PermalinkForBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
