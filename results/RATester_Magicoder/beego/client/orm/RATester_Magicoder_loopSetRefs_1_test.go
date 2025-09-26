package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestloopSetRefs_1(t *testing.T) {
	type args struct {
		refs     []interface{}
		sInds    []reflect.Value
		nIndsPtr *[]reflect.Value
		eTyps    []reflect.Type
		init     bool
	}
	tests := []struct {
		name string
		args args
		want []reflect.Value
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

			o := &rawSet{}
			o.loopSetRefs(tt.args.refs, tt.args.sInds, tt.args.nIndsPtr, tt.args.eTyps, tt.args.init)
		})
	}
}
