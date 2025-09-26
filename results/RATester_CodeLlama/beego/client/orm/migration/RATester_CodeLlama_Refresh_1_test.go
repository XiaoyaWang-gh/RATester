package migration

import (
	"fmt"
	"testing"
)

func TestRefresh_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err := Refresh()
	if err != nil {
		t.Error("Refresh error:", err)
	}
}
