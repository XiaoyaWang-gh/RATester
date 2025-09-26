package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextNonSpace_1(t *testing.T) {
	tests := []struct {
		name string
		t    *Tree
		want item
	}{
		{
			name: "test1",
			t: &Tree{
				token: [3]item{
					{typ: itemSpace, val: " "},
					{typ: itemSpace, val: " "},
					{typ: itemSpace, val: " "},
				},
			},
			want: item{typ: itemSpace, val: " "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.nextNonSpace(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.nextNonSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
