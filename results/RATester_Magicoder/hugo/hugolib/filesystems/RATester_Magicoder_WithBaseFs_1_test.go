package filesystems

import (
	"fmt"
	"testing"
)

func TestWithBaseFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BaseFs{
		theBigFs:          &filesystemsCollector{},
		SourceFilesystems: &SourceFilesystems{},
	}
	bb := &BaseFs{}

	err := WithBaseFs(b)(bb)
	if err != nil {
		t.Errorf("WithBaseFs returned an error: %v", err)
	}

	if bb.theBigFs != b.theBigFs {
		t.Errorf("WithBaseFs did not set theBigFs correctly")
	}
	if bb.SourceFilesystems != b.SourceFilesystems {
		t.Errorf("WithBaseFs did not set SourceFilesystems correctly")
	}
}
