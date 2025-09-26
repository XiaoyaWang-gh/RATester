package web

import (
	"fmt"
	"testing"
)

func TestRegisterAdmin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if err := registerAdmin(); err != nil {
		t.Errorf("registerAdmin() error = %v", err)
		return
	}
}
