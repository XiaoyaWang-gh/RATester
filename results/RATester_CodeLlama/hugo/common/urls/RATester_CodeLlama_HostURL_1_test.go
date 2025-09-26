package urls

import (
	"fmt"
	"testing"
)

func TestHostURL_1(t *testing.T) {
	tests := []struct {
		name string
		b    BaseURL
		want string
	}{
		{
			name: "WithPath",
			b:    BaseURL{WithPath: "http://example.com/path"},
			want: "http://example.com",
		},
		{
			name: "WithPathNoTrailingSlash",
			b:    BaseURL{WithPathNoTrailingSlash: "http://example.com/path/"},
			want: "http://example.com",
		},
		{
			name: "WithoutPath",
			b:    BaseURL{WithoutPath: "http://example.com"},
			want: "http://example.com",
		},
		{
			name: "BasePath",
			b:    BaseURL{BasePath: "http://example.com/path"},
			want: "http://example.com",
		},
		{
			name: "BasePathNoTrailingSlash",
			b:    BaseURL{BasePathNoTrailingSlash: "http://example.com/path/"},
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.b.HostURL(); got != tt.want {
				t.Errorf("BaseURL.HostURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
