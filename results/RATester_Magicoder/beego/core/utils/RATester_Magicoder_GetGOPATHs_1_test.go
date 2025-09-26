package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestGetGOPATHs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("GOPATH", "/usr/local/go")
	defer os.Unsetenv("GOPATH")

	gopaths := GetGOPATHs()

	if len(gopaths) != 1 || gopaths[0] != "/usr/local/go" {
		t.Errorf("Expected GOPATHs to be ['/usr/local/go'], but got %v", gopaths)
	}
}
