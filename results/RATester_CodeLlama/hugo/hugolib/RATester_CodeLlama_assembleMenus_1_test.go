package hugolib

import (
	"fmt"
	"testing"
)

func TestAssembleMenus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.assembleMenus()
}
