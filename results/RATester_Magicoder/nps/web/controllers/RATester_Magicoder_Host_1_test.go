package controllers

import (
	"fmt"
	"testing"
)

func TestHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.Host()

	// Add assertions here to verify the expected behavior
}
