package env

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestEnvFile_1(t *testing.T) {
	t.Run("GOENV is set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOENV", "test")
		defer os.Unsetenv("GOENV")

		file, err := envFile()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if file != "test" {
			t.Errorf("Expected file to be 'test', got %s", file)
		}
	})

	t.Run("GOENV is off", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOENV", "off")
		defer os.Unsetenv("GOENV")

		_, err := envFile()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != "GOENV=off" {
			t.Errorf("Expected error to be 'GOENV=off', got %s", err.Error())
		}
	})

	t.Run("GOENV is not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("GOENV")

		file, err := envFile()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if file != filepath.Join(os.Getenv("HOME"), ".go", "env") {
			t.Errorf("Expected file to be '%s', got %s", filepath.Join(os.Getenv("HOME"), ".go", "env"), file)
		}
	})
}
