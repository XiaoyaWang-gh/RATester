package releaser

import (
	"fmt"
	"testing"
)

func TestreplaceInFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &ReleaseHandler{
		branchVersion: "0.90.0",
		step:          1,
		skipPush:      true,
		try:           true,
		git:           func(args ...string) (string, error) { return "", nil },
	}

	err := handler.replaceInFile("testfile.txt", "old", "new")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
