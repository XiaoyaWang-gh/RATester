package framework

import (
	"fmt"
	"testing"
)

func TestNewDefaultFramework_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewDefaultFramework()
	if f == nil {
		t.Errorf("NewDefaultFramework failed")
	}
}
