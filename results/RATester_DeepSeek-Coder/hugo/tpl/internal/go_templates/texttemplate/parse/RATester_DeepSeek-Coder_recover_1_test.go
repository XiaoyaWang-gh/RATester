package parse

import (
	"fmt"
	"testing"
)

func TestRecover_1(t *testing.T) {
	type args struct {
		errp *error
	}
	tests := []struct {
		name string
		t    *Tree
		args args
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

			tt.t.recover(tt.args.errp)
		})
	}
}
