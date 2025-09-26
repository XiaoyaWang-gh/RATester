package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTextOrAction_1(t *testing.T) {
	tests := []struct {
		name string
		t    *Tree
		want Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.textOrAction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.textOrAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
