package logs

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestcreateLogFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
		Perm:     "0666",
	}

	fd, err := w.createLogFile()
	if err != nil {
		t.Errorf("createLogFile() error = %v", err)
		return
	}

	defer fd.Close()

	fi, err := fd.Stat()
	if err != nil {
		t.Errorf("Stat() error = %v", err)
		return
	}

	perm, err := strconv.ParseInt(w.Perm, 8, 64)
	if err != nil {
		t.Errorf("ParseInt() error = %v", err)
		return
	}

	if fi.Mode().Perm() != os.FileMode(perm) {
		t.Errorf("createLogFile() file permission = %v, want %v", fi.Mode().Perm(), os.FileMode(perm))
	}
}
