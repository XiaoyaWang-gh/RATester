package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPipeline_1(t *testing.T) {
	type args struct {
		context string
		end     itemType
	}
	tests := []struct {
		name     string
		t        *Tree
		args     args
		wantPipe *PipeNode
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

			if gotPipe := tt.t.pipeline(tt.args.context, tt.args.end); !reflect.DeepEqual(gotPipe, tt.wantPipe) {
				t.Errorf("Tree.pipeline() = %v, want %v", gotPipe, tt.wantPipe)
			}
		})
	}
}
