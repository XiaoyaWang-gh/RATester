package paths

import (
	"fmt"
	"testing"
)

func TestBase_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pathBridge := pathBridge{}
	in := "path/to/file.txt"
	expected := "file.txt"
	actual := pathBridge.Base(in)
	if actual != expected {
		t.Errorf("Base(%q) = %q; want %q", in, actual, expected)
	}
}
