package template

import (
	"fmt"
	"testing"
)

func TestjsStrEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"test", contentTypeJSStr},
			want: "test",
		},
		{
			name: "Test case 2",
			args: []any{"test", contentTypeJSStr},
			want: "test",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := jsStrEscaper(tt.args...); got != tt.want {
				t.Errorf("jsStrEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
