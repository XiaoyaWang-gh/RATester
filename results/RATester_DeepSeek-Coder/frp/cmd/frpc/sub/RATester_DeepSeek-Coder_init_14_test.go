package sub

import (
	"fmt"
	"testing"
)

func TestInit_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootCmd.AddCommand(verifyCmd)

	if len(rootCmd.Commands()) == 0 {
		t.Errorf("Expected rootCmd to have commands, but it does not")
	}
}
