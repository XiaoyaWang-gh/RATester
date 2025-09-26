package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRange_1(t *testing.T) {
	type args struct {
		pos      Pos
		line     int
		pipe     *PipeNode
		list     *ListNode
		elseList *ListNode
	}
	tests := []struct {
		name string
		args args
		want *RangeNode
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

			if got := (&Tree{}).newRange(tt.args.pos, tt.args.line, tt.args.pipe, tt.args.list, tt.args.elseList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
