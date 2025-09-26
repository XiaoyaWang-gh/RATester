package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewAction_1(t *testing.T) {
	type args struct {
		pos  Pos
		line int
		pipe *PipeNode
	}
	tests := []struct {
		name string
		args args
		want *ActionNode
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

			tr := &Tree{}
			if got := tr.newAction(tt.args.pos, tt.args.line, tt.args.pipe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.newAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
