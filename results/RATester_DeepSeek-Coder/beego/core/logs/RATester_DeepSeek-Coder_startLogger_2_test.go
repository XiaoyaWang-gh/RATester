package logs

import (
	"fmt"
	"testing"
)

func TestStartLogger_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
	}

	err := w.startLogger()
	if err != nil {
		t.Errorf("startLogger() error = %v", err)
		return
	}

	if w.fileWriter == nil {
		t.Errorf("startLogger() error = fileWriter is nil")
	}

	w.fileWriter.Close()
}
