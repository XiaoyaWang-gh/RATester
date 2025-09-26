package web

import (
	"fmt"
	"testing"
)

func TestInclude_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	cList := []ControllerInterface{}
	n.Include(cList...)
}
