package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBreakControl_1(t *testing.T) {
	type args struct {
		pos  Pos
		line int
	}
	tests := []struct {
		name string
		t    *Tree
		args args
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

			if got := tt.t.breakControl(tt.args.pos, tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.breakControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
