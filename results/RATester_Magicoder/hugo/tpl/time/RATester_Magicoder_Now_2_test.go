package time

import (
	"fmt"
	"testing"
)

func TestNow_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	now := ns.Now()

	if now.IsZero() {
		t.Error("Expected Now() to return a non-zero time")
	}
}
