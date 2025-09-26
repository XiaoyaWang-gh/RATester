package math

import (
	"fmt"
	"testing"
)

func TestRand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	if ns.Rand() == ns.Rand() {
		t.Errorf("Rand() should return a different value each time")
	}
}
