package htesting

import (
	"fmt"
	"testing"
)

func TestRandBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	for i := 0; i < 1000; i++ {
		if RandBool() != (Rnd.Intn(2) != 0) {
			t.Errorf("RandBool() = %v, want %v", RandBool(), Rnd.Intn(2) != 0)
		}
	}
}
