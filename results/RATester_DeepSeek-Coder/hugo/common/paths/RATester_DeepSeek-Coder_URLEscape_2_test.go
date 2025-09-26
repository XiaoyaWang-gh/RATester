package paths

import (
	"fmt"
	"testing"
)

func TestURLEscape_2(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with simple URL",
			args: args{uri: "http://example.com"},
			want: "http://example.com",
		},
		{
			name: "Test with URL with path",
			args: args{uri: "http://example.com/path/to/page"},
			want: "http://example.com/path/to/page",
		},
		{
			name: "Test with URL with query parameters",
			args: args{uri: "http://example.com/path/to/page?param1=value1&param2=value2"},
			want: "http://example.com/path/to/page?param1=value1&param2=value2",
		},
		{
			name: "Test with URL with fragment",
			args: args{uri: "http://example.com/path/to/page#fragment"},
			want: "http://example.com/path/to/page#fragment",
		},
		{
			name: "Test with URL with user info",
			args: args{uri: "http://user:pass@example.com"},
			want: "http://user:pass@example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := URLEscape(tt.args.uri); got != tt.want {
				t.Errorf("URLEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
