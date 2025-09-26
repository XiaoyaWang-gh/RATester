package output

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeLayoutsPresentable_1(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want []string
	}{
		{
			name: "Test 1",
			arg:  []string{"_text/indexes/index.md", "page/index.md", "indexes/index.md", "index.md"},
			want: []string{"layouts/index.md"},
		},
		{
			name: "Test 2",
			arg:  []string{"_text/indexes/index.md", "page/index.md", "indexes/index.md", "index.md", "indexes/indexes.md"},
			want: []string{"layouts/index.md", "layouts/indexes.md"},
		},
		{
			name: "Test 3",
			arg:  []string{"_text/indexes/index.md", "page/index.md", "indexes/index.md", "index.md", "indexes/indexes.md", "indexes/indexes/index.md"},
			want: []string{"layouts/index.md", "layouts/indexes.md"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := makeLayoutsPresentable(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeLayoutsPresentable() = %v, want %v", got, tt.want)
			}
		})
	}
}
