package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestescapeSpecialScriptTags_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "Test case 1",
			s:    []byte("<script>alert('Hello, world!')</script>"),
			want: []byte("&lt;script&gt;alert('Hello, world!')&lt;/script&gt;"),
		},
		{
			name: "Test case 2",
			s:    []byte("<script>alert('Hello, world!')</script>"),
			want: []byte("&lt;script&gt;alert('Hello, world!')&lt;/script&gt;"),
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

			if got := escapeSpecialScriptTags(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("escapeSpecialScriptTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
