package file

import (
	"fmt"
	"testing"
)

func TestLess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := PairList{
		{
			clientFlow: &Flow{
				ExportFlow: 10,
				InletFlow:  10,
				FlowLimit:  10,
			},
			key:   "ExportFlow",
			order: "desc",
		},
		{
			clientFlow: &Flow{
				ExportFlow: 10,
				InletFlow:  10,
				FlowLimit:  10,
			},
			key:   "InletFlow",
			order: "desc",
		},
		{
			clientFlow: &Flow{
				ExportFlow: 10,
				InletFlow:  10,
				FlowLimit:  10,
			},
			key:   "FlowLimit",
			order: "desc",
		},
	}
	if !p.Less(0, 1) {
		t.Errorf("Less(0, 1) = %v, want true", p.Less(0, 1))
	}
	if p.Less(0, 2) {
		t.Errorf("Less(0, 2) = %v, want false", p.Less(0, 2))
	}
	if p.Less(1, 2) {
		t.Errorf("Less(1, 2) = %v, want false", p.Less(1, 2))
	}
	if !p.Less(2, 1) {
		t.Errorf("Less(2, 1) = %v, want true", p.Less(2, 1))
	}
}
