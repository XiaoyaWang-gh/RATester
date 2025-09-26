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
			name: "Test case 1",
			els:  []string{"test1", "test2"},
			want: []string{"test1", "test2", "test3", "test4"},
		},
		{
			name: "Test case 2",
			els:  []string{"test5", "test6"},
			want: []string{"test5", "test6", "test7", "test8"},
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
			p.Prepend("test3", "test4")
			if !reflect.DeepEqual(p.els, tt.want) {
				t.Errorf("Prepend() = %v, want %v", p.els, tt.want)
			}
		})
	}
}
