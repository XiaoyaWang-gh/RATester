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

	// Add assertions here to verify the expected behavior
}
