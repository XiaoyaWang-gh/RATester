package maps

import (
	"fmt"
	"testing"
)

func Testmerge_3(t *testing.T) {
	type args struct {
		ps ParamsMergeStrategy
		pp Params
	}
	tests := []struct {
		name string
		p    Params
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

			tt.p.merge(tt.args.ps, tt.args.pp)
		})
	}
}
