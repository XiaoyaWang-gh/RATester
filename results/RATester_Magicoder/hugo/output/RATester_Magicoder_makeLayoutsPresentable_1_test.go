package output

import (
	"fmt"
	"reflect"
	"testing"
)

func TestmakeLayoutsPresentable_1(t *testing.T) {
	tests := []struct {
		name string
		l    []string
		want []string
	}{
		{
			name: "Test case 1",
			l:    []string{"page/test.md", "_text/test2.md", "test3.md"},
			want: []string{"layouts/test3.md"},
		},
		{
			name: "Test case 2",
			l:    []string{"page/test.md", "_text/test2.md", "test3.md", "indexes/test4.md"},
			want: []string{"layouts/test3.md"},
		},
		{
			name: "Test case 3",
			l:    []string{"page/test.md", "_text/test2.md", "test3.md", "indexes/test4.md", "test5.md"},
			want: []string{"layouts/test3.md", "layouts/test5.md"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeLayoutsPresentable(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeLayoutsPresentable() = %v, want %v", got, tt.want)
			}
		})
	}
}
