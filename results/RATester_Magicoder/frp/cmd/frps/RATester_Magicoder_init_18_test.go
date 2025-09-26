package main

import (
	"fmt"
	"testing"
)

func TestInit_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootCmd.AddCommand(verifyCmd)

	// Add assertions here to verify the expected behavior
}
