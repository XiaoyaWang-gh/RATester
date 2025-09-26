package collections

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestLookupFunc_1(t *testing.T) {
	type args struct {
		ctx   context.Context
		fname string
	}
	tests := []struct {
		name  string
		ns    *Namespace
		args  args
		want  reflect.Value
		want1 bool
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

			got, got1 := tt.ns.lookupFunc(tt.args.ctx, tt.args.fname)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookupFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("lookupFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
