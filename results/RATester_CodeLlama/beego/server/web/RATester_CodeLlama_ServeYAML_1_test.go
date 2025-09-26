package web

import (
	"fmt"
	"testing"
)

func TestServeYAML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	c := &Controller{}
	err := c.ServeYAML()
	if err != nil {
		t.Errorf("ServeYAML() error = %v", err)
		return
	}
	// TODO: verify your expectations
}
