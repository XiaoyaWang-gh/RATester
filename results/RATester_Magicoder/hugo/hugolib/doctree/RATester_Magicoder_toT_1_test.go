package doctree

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoT_1(t *testing.T) {
	type args struct {
		tree *NodeShiftTree[any]
		v    any
	}
	tests := []struct {
		name      string
		r         *NodeShiftTreeWalker[any]
		args      args
		wantT     any
		wantOk    bool
		wantExact DimensionFlag
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

			gotT, gotOk, gotExact := tt.r.toT(tt.args.tree, tt.args.v)
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("toT() gotT = %v, want %v", gotT, tt.wantT)
			}
			if gotOk != tt.wantOk {
				t.Errorf("toT() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotExact != tt.wantExact {
				t.Errorf("toT() gotExact = %v, want %v", gotExact, tt.wantExact)
			}
		})
	}
}
