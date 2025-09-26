package client

import (
	"fmt"
	"testing"
)

func TestPreHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	c := &core{}
	c.client = &Client{}
	c.req = &Request{}

	// Act
	err := c.preHooks()

	// Assert
	if err != nil {
		t.Errorf("preHooks() error = %v, want nil", err)
		return
	}
}
