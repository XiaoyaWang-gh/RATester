package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeParamsWithStrategy_1(t *testing.T) {
	type args struct {
		strategy string
		dst      Params
		src      Params
	}
	tests := []struct {
		name string
		args args
		want Params
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

			MergeParamsWithStrategy(tt.args.strategy, tt.args.dst, tt.args.src)
			if !reflect.DeepEqual(tt.args.dst, tt.want) {
				t.Errorf("MergeParamsWithStrategy() = %v, want %v", tt.args.dst, tt.want)
			}
		})
	}
}
