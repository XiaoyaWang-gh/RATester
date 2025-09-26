package hugo

import (
	"fmt"
	"testing"
)

func TestFormatDep_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "github.com/golang/dep"
	version := "v0.3.0"
	want := "github.com/golang/dep=v0.3.0"
	if got := formatDep(path, version); got != want {
		t.Errorf("formatDep(%q, %q) = %q; want %q", path, version, got, want)
	}
}
