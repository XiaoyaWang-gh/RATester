package captcha

import (
	"fmt"
	"testing"
)

func TestRandFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	from := 1.0
	to := 10.0
	got := randFloat(from, to)
	if got < from || got > to {
		t.Errorf("randFloat(%f, %f) = %f, want in [%f, %f]", from, to, got, from, to)
	}
}
