package paths

import (
	"fmt"
	"testing"
)

func TestDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pathBridge := pathBridge{}
	in := "path/to/file.txt"
	expected := "path/to"
	actual := pathBridge.Dir(in)
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
