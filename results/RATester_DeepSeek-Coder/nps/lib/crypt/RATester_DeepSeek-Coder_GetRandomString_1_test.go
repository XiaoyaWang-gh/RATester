package crypt

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetRandomString_1(t *testing.T) {
	t.Run("Testing with length 5", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := GetRandomString(5)
		if len(result) != 5 {
			t.Errorf("Expected length of 5, got %d", len(result))
		}
		for _, char := range result {
			if !strings.ContainsRune("0123456789abcdefghijklmnopqrstuvwxyz", char) {
				t.Errorf("Expected only characters from '0123456789abcdefghijklmnopqrstuvwxyz', got %c", char)
			}
		}
	})

	t.Run("Testing with length 10", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := GetRandomString(10)
		if len(result) != 10 {
			t.Errorf("Expected length of 10, got %d", len(result))
		}
		for _, char := range result {
			if !strings.ContainsRune("0123456789abcdefghijklmnopqrstuvwxyz", char) {
				t.Errorf("Expected only characters from '0123456789abcdefghijklmnopqrstuvwxyz', got %c", char)
			}
		}
	})
}
