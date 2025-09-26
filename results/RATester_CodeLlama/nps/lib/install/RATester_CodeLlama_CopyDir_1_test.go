package install

import (
	"fmt"
	"os"
	"testing"

	"ehang.io/nps/lib/common"
)

func TestCopyDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srcPath := "./testdata/src"
	destPath := "./testdata/dest"
	err := CopyDir(srcPath, destPath)
	if err != nil {
		t.Errorf("CopyDir failed, err: %v", err)
	}
	if !common.IsWindows() {
		if err := os.RemoveAll(destPath); err != nil {
			t.Errorf("RemoveAll failed, err: %v", err)
		}
	}
}
