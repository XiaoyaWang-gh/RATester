package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseControl_1(t *testing.T) {
	type args struct {
		context string
	}
	tests := []struct {
		name         string
		t            *Tree
		args         args
		wantPos      Pos
		wantLine     int
		wantPipe     *PipeNode
		wantList     *ListNode
		wantElseList *ListNode
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

			gotPos, gotLine, gotPipe, gotList, gotElseList := tt.t.parseControl(tt.args.context)
			if gotPos != tt.wantPos {
				t.Errorf("Tree.parseControl() gotPos = %v, want %v", gotPos, tt.wantPos)
			}
			if gotLine != tt.wantLine {
				t.Errorf("Tree.parseControl() gotLine = %v, want %v", gotLine, tt.wantLine)
			}
			if !reflect.DeepEqual(gotPipe, tt.wantPipe) {
				t.Errorf("Tree.parseControl() gotPipe = %v, want %v", gotPipe, tt.wantPipe)
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("Tree.parseControl() gotList = %v, want %v", gotList, tt.wantList)
			}
			if !reflect.DeepEqual(gotElseList, tt.wantElseList) {
				t.Errorf("Tree.parseControl() gotElseList = %v, want %v", gotElseList, tt.wantElseList)
			}
		})
	}
}
