package media

import (
	"fmt"
	"testing"
)

func TestSwap_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	types := Types{
		{Type: "type1"},
		{Type: "type2"},
	}

	types.Swap(0, 1)

	if types[0].Type != "type2" || types[1].Type != "type1" {
		t.Errorf("Swap failed. Expected: [type2, type1], Got: %v", types)
	}
}
