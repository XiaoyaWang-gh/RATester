package framework

import (
	"fmt"
	"testing"
)

func TestAfterEach_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewDefaultFramework()
	f.AfterEach()
}
