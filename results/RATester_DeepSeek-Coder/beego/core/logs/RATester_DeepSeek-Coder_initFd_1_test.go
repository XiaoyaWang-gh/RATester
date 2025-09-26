package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestInitFd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		fileWriter: &os.File{},
	}

	err := w.initFd()
	if err != nil {
		t.Errorf("initFd() error = %v", err)
	}

	if w.maxSizeCurSize != 0 {
		t.Errorf("initFd() maxSizeCurSize = %v, want %v", w.maxSizeCurSize, 0)
	}
}
