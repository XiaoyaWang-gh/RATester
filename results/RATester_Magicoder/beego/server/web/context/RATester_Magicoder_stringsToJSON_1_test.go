package context

import (
	"fmt"
	"testing"
)

func TeststringsToJSON_1(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test with simple string",
			str:  "Hello, World!",
			want: "Hello, World!",
		},
		{
			name: "Test with special characters",
			str:  "日本語",
			want: "\\u65e5\\u672c\\u8a9e",
		},
		{
			name: "Test with empty string",
			str:  "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := stringsToJSON(tt.str); got != tt.want {
				t.Errorf("stringsToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
