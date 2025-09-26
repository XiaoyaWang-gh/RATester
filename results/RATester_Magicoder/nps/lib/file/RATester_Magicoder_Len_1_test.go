package file

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pairList := PairList{
		&Pair{key: "key1", cId: 1, order: "order1", clientFlow: &Flow{ExportFlow: 10, InletFlow: 20, FlowLimit: 30}},
		&Pair{key: "key2", cId: 2, order: "order2", clientFlow: &Flow{ExportFlow: 10, InletFlow: 20, FlowLimit: 30}},
		&Pair{key: "key3", cId: 3, order: "order3", clientFlow: &Flow{ExportFlow: 10, InletFlow: 20, FlowLimit: 30}},
	}

	expected := 3
	actual := pairList.Len()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
