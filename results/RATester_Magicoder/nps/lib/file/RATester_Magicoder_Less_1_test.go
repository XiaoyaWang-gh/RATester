package file

import (
	"fmt"
	"testing"
)

func TestLess_1(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    PairList
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

			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("PairList.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
