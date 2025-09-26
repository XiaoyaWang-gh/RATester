package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewParam_1(t *testing.T) {
	type args struct {
		name   string
		parser paramParser
		opts   []MethodParamOption
	}
	tests := []struct {
		name      string
		args      args
		wantParam *MethodParam
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

			if gotParam := newParam(tt.args.name, tt.args.parser, tt.args.opts); !reflect.DeepEqual(gotParam, tt.wantParam) {
				t.Errorf("newParam() = %v, want %v", gotParam, tt.wantParam)
			}
		})
	}
}
