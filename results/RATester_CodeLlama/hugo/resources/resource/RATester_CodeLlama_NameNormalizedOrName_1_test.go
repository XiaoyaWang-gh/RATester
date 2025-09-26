package resource

import (
	"fmt"
	"testing"
)

func TestNameNormalizedOrName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var r Resource
	if got := NameNormalizedOrName(r); got != "" {
		t.Errorf("NameNormalizedOrName() = %v, want %v", got, "")
	}
}
