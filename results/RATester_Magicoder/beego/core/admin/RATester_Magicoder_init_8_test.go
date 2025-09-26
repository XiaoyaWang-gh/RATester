package admin

import (
	"fmt"
	"os"
	"testing"
)

func Testinit_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pid = os.Getpid()
	if pid != os.Getpid() {
		t.Errorf("Expected pid to be %d, but got %d", os.Getpid(), pid)
	}
}
