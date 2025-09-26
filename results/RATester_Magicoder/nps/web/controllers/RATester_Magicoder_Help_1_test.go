package controllers

import (
	"fmt"
	"testing"
)

func TestHelp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.Help()

	// Add assertions here to validate the behavior of the method under test
}
