package doctree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSkipPrefix_1(t *testing.T) {
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
	type args struct {
		prefix []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
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

			r := &NodeShiftTreeWalker[int]{
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
			r.SkipPrefix(tt.args.prefix...)
			if !reflect.DeepEqual(r.skipPrefixes, tt.want) {
				t.Errorf("SkipPrefix() = %v, want %v", r.skipPrefixes, tt.want)
			}
		})
	}
}
