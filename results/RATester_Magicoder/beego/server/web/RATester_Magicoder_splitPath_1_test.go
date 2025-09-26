package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsplitPath_1(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want []string
	}{
		{
			name: "empty string",
			key:  "",
			want: []string{},
		},
		{
			name: "single element",
			key:  "element",
			want: []string{"element"},
		},
		{
			name: "multiple elements",
			key:  "element1/element2/element3",
			want: []string{"element1", "element2", "element3"},
		},
		{
			name: "leading and trailing slashes",
			key:  "/element1/element2/element3/",
			want: []string{"element1", "element2", "element3"},
		},
		{
			name: "leading and trailing spaces",
			key:  " element1/element2/element3 ",
			want: []string{"element1", "element2", "element3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := splitPath(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
