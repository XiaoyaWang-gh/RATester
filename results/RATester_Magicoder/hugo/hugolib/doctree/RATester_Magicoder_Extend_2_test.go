package doctree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtend_2(t *testing.T) {
	type fields struct {
		Tree         *NodeShiftTree[int]
		Handle       func(s string, v int, exact DimensionFlag) (terminate bool, err error)
		Prefix       string
		LockType     LockType
		NoShift      bool
		Exact        bool
		Debug        bool
		WalkContext  *WalkContext[int]
		skipPrefixes []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *NodeShiftTreeWalker[int]
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

			r := NodeShiftTreeWalker[int]{
				Tree:         tt.fields.Tree,
				Handle:       tt.fields.Handle,
				Prefix:       tt.fields.Prefix,
				LockType:     tt.fields.LockType,
				NoShift:      tt.fields.NoShift,
				Exact:        tt.fields.Exact,
				Debug:        tt.fields.Debug,
				WalkContext:  tt.fields.WalkContext,
				skipPrefixes: tt.fields.skipPrefixes,
			}
			if got := r.Extend(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeShiftTreeWalker.Extend() = %v, want %v", got, tt.want)
			}
		})
	}
}
