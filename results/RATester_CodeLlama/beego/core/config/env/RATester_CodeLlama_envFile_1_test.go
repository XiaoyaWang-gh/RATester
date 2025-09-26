package env

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvFile_1(t *testing.T) {
	t.Run("GOENV=off", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOENV", "off")
		defer os.Unsetenv("GOENV")
		_, err := envFile()
		if err == nil {
			t.Error("expected error")
		}
	})
	t.Run("GOENV=on", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOENV", "on")
		defer os.Unsetenv("GOENV")
		_, err := envFile()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("GOENV=", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOENV", "")
		defer os.Unsetenv("GOENV")
		_, err := envFile()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("GOENV=missing", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("GOENV")
		_, err := envFile()
		if err != nil {
			t.Error(err)
		}
	})
}
