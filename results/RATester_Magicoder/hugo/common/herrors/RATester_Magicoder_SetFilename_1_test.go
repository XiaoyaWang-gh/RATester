package herrors

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestSetFilename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fe := &fileError{
		position: text.Position{
			Filename: "test.txt",
		},
	}

	newFilename := "new_test.txt"
	fe.SetFilename(newFilename)

	if fe.position.Filename != newFilename {
		t.Errorf("Expected filename to be %s, but got %s", newFilename, fe.position.Filename)
	}
}
