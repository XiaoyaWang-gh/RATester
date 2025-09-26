package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOperand_1(t *testing.T) {
	tests := []struct {
		name string
		t    *Tree
		want Node
	}{
		{
			name: "test1",
			t:    &Tree{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.operand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operand() = %v, want %v", got, tt.want)
			}
		})
	}
}
