package cache

import (
	"fmt"
	"testing"
)

func TestDefaultExpiredFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	got := defaultExpiredFunc()
	if got == nil {
		t.Fatal("got nil")
	}
	if got() == 0 {
		t.Error("got 0")
	}
}
