package server

import (
	"fmt"
	"testing"
)

func TestReplaced_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	newCtl := &Control{}
	ctl.Replaced(newCtl)
}
