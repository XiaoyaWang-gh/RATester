package pagemeta

import (
	"fmt"
	"testing"
)

func TestIsZero_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := BuildConfig{}
	if !b.IsZero() {
		t.Errorf("BuildConfig is zero")
	}
}
