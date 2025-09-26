package related

import (
	"fmt"
	"testing"
)

func TestSwap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := ranks{&rank{}, &rank{}}
	r[0].addWeight(10)
	r[1].addWeight(20)
	r.Swap(0, 1)
	if r[0].Weight != 20 || r[1].Weight != 10 {
		t.Errorf("Swap failed. Expected: %d, %d. Got: %d, %d", 20, 10, r[0].Weight, r[1].Weight)
	}
}
