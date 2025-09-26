package file

import (
	"fmt"
	"testing"
)

func TestSwap_1(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    PairList
		args args
	}{
		{
			name: "TestSwap",
			p:    PairList{&Pair{key: "key1", cId: 1, order: "order1"}, &Pair{key: "key2", cId: 2, order: "order2"}},
			args: args{i: 0, j: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.p.Swap(tt.args.i, tt.args.j)
			if tt.p[0].key != "key2" || tt.p[1].key != "key1" {
				t.Errorf("Expected PairList to be swapped, but got %v", tt.p)
			}
		})
	}
}
