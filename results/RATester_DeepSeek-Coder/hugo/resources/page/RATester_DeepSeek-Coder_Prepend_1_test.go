package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrepend_1(t *testing.T) {
	tests := []struct {
		name string
		els  []string
		want []string
	}{
		{
			name: "Test with one element",
			els:  []string{"test1"},
			want: []string{"test1"},
		},
		{
			name: "Test with multiple elements",
			els:  []string{"test1", "test2", "test3"},
			want: []string{"test1", "test2", "test3"},
		},
		{
			name: "Test with no elements",
			els:  []string{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &pagePathBuilder{
				els: tt.els,
			}
			p.Prepend(tt.els...)
			if !reflect.DeepEqual(p.els, tt.want) {
				t.Errorf("got %v, want %v", p.els, tt.want)
			}
		})
	}
}
