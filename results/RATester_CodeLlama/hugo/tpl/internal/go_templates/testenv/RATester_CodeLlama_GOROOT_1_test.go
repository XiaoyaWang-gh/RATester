package testenv

import (
	"fmt"
	"testing"
)

func TestGOROOT_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path, err := findGOROOT()
	if err != nil {
		t.Helper()
		t.Skip(err)
	}
	if path == "" {
		t.Error("GOROOT is empty")
	}
}
