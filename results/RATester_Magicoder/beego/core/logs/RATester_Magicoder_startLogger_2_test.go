package logs

import (
	"fmt"
	"testing"
)

func TeststartLogger_2(t *testing.T) {
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
		t.Errorf("startLogger() failed with %s", err)
	}

	if w.fileWriter == nil {
		t.Error("startLogger() did not open the file")
	}

	w.fileWriter.Close()
}
