package system

import (
	"fmt"
	"testing"
)

func TestEnableCompatibilityMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	EnableCompatibilityMode()

	// Add your assertions here
}
