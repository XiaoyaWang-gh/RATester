package controllers

import (
	"fmt"
	"testing"
)

func Testerror_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &BaseController{}
	ctrl.error()

	// Add assertions here
}
