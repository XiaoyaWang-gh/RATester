package identity

import (
	"fmt"
	"testing"
)

func TestWalkIdentitiesShallow_1(t *testing.T) {
	type args struct {
		v     any
		level int
		cb    func(level int, id Identity) bool
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := walkIdentitiesShallow(tt.args.v, tt.args.level, tt.args.cb); got != tt.want {
				t.Errorf("walkIdentitiesShallow() = %v, want %v", got, tt.want)
			}
		})
	}
}
