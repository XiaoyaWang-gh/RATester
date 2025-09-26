package file

import (
	"fmt"
	"testing"
)

func TestSwap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pair1 := &Pair{key: "key1", cId: 1, order: "order1"}
	pair2 := &Pair{key: "key2", cId: 2, order: "order2"}

	pairList := PairList{pair1, pair2}

	pairList.Swap(0, 1)

	if pairList[0] != pair2 || pairList[1] != pair1 {
		t.Errorf("Swap failed. Expected: %v, %v. Got: %v, %v", pair2, pair1, pairList[0], pairList[1])
	}
}
