package controllers

import (
	"fmt"
	"testing"
)

func TestAjax_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
