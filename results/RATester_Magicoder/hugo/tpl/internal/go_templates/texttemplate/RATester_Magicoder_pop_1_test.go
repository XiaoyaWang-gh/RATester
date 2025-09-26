package template

import (
	"fmt"
	"testing"
)

func Testpop_1(t *testing.T) {
	type args struct {
		mark int
	}
	tests := []struct {
		name string
		s    *state
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

			tt.s.pop(tt.args.mark)
		})
	}
}
