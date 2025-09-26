package web

import (
	"fmt"
	"testing"
)

func TestregisterAdmin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	BConfig.Listen.EnableAdmin = true
	err := registerAdmin()
	if err != nil {
		t.Errorf("registerAdmin() failed, expected nil, got %v", err)
	}
}
