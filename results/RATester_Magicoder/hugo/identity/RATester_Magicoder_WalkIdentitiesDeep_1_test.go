package identity

import (
	"fmt"
	"testing"
)

func TestWalkIdentitiesDeep_1(t *testing.T) {
	type args struct {
		v  any
		cb func(level int, id Identity) bool
	}
	tests := []struct {
		name string
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

			WalkIdentitiesDeep(tt.args.v, tt.args.cb)
		})
	}
}
