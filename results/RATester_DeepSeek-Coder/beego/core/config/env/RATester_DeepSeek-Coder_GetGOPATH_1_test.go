package env

import (
	"fmt"
	"go/build"
	"os"
	"testing"
)

func TestGetGOPATH_1(t *testing.T) {

	t.Run("Test when GOPATH is set in environment", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOPATH", "/home/user/go")
		defer os.Unsetenv("GOPATH")
		got := GetGOPATH()
		want := "/home/user/go"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Test when GOPATH is not set in environment", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("GOPATH")
		got := GetGOPATH()
		want := build.Default.GOPATH
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Test when GOPATH is empty string in environment", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GOPATH", "")
		defer os.Unsetenv("GOPATH")
		got := GetGOPATH()
		want := build.Default.GOPATH
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
