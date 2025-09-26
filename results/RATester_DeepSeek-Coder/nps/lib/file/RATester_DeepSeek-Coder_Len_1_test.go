package file

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	tests := []struct {
		name string
		p    PairList
		want int
	}{
		{
			name: "Test case 1",
			p:    PairList{&Pair{key: "key1", cId: 1, order: "order1", clientFlow: &Flow{ExportFlow: 1, InletFlow: 1, FlowLimit: 1}}},
			want: 1,
		},
		{
			name: "Test case 2",
			p:    PairList{&Pair{key: "key1", cId: 1, order: "order1", clientFlow: &Flow{ExportFlow: 1, InletFlow: 1, FlowLimit: 1}}, &Pair{key: "key2", cId: 2, order: "order2", clientFlow: &Flow{ExportFlow: 1, InletFlow: 1, FlowLimit: 1}}},
			want: 2,
		},
		{
			name: "Test case 3",
			p:    PairList{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.Len(); got != tt.want {
				t.Errorf("PairList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
