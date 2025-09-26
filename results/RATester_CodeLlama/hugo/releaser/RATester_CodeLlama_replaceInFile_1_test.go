package releaser

import (
	"fmt"
	"os"
	"testing"
)

func TestReplaceInFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &ReleaseHandler{}
	filename := "testdata/test.txt"
	oldNew := []string{"foo", "bar"}
	err := r.replaceInFile(filename, oldNew...)
	if err != nil {
		t.Fatal(err)
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "bar" {
		t.Errorf("Expected %q, got %q", "bar", string(b))
	}
}
