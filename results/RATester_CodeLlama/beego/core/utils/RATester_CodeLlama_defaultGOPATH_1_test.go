package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDefaultGOPATH_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if runtime.GOOS == "windows" {
		if defaultGOPATH() != filepath.Join(os.Getenv("USERPROFILE"), "go") {
			t.Errorf("defaultGOPATH() = %v, want %v", defaultGOPATH(), filepath.Join(os.Getenv("USERPROFILE"), "go"))
		}
	} else if runtime.GOOS == "plan9" {
		if defaultGOPATH() != filepath.Join(os.Getenv("home"), "go") {
			t.Errorf("defaultGOPATH() = %v, want %v", defaultGOPATH(), filepath.Join(os.Getenv("home"), "go"))
		}
	} else {
		if defaultGOPATH() != filepath.Join(os.Getenv("HOME"), "go") {
			t.Errorf("defaultGOPATH() = %v, want %v", defaultGOPATH(), filepath.Join(os.Getenv("HOME"), "go"))
		}
	}
}
