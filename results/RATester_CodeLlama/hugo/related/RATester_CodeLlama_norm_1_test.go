package related

import (
	"fmt"
	"testing"
)

func TestNorm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	num := 10
	min := 1
	max := 10
	want := 10
	got := norm(num, min, max)
	if got != want {
		t.Errorf("norm(%d, %d, %d) = %d; want %d", num, min, max, got, want)
	}
}
