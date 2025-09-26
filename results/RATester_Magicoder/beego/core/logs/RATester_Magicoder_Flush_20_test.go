package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestFlush_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		fileWriter: &os.File{},
	}

	w.Flush()

	if w.fileWriter.Sync() != nil {
		t.Errorf("Expected Flush to call Sync on the fileWriter, but it didn't")
	}
}
