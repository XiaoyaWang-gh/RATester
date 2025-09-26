package orm

import (
	"fmt"
	"testing"
)

func TestNewOrm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := NewOrm()
	if orm == nil {
		t.Error("Expected non-nil orm, got nil")
	}
}
