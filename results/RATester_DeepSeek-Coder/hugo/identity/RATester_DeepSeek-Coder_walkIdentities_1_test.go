package identity

import (
	"fmt"
	"testing"
)

func TestWalkIdentities_1(t *testing.T) {
	type args struct {
		v     any
		level int
		deep  bool
		seen  map[Identity]bool
		cb    func(level int, id Identity) bool
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

			walkIdentities(tt.args.v, tt.args.level, tt.args.deep, tt.args.seen, tt.args.cb)
		})
	}
}
