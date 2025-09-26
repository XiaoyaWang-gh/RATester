package web

import (
	"fmt"
	"testing"
)

func TestFinish_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		// initialize your controller here
	}

	// call the method under test
	ctrl.Finish()

	// add assertions here
}
