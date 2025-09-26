package controllers

import (
	"fmt"
	"testing"
)

func TestHttp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.Http()
	// Add your assertions here
}
