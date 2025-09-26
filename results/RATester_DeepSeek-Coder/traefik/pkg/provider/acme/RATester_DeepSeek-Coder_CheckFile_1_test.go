package acme

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = CheckFile(tmpFile.Name())
	if err != nil {
		t.Errorf("CheckFile() error = %v, wantErr %v", err, false)
		return
	}

	_, err = CheckFile("non-existing-file")
	if err == nil {
		t.Errorf("CheckFile() error = %v, wantErr %v", err, true)
		return
	}

	_, err = CheckFile("/dev/null")
	if err == nil {
		t.Errorf("CheckFile() error = %v, wantErr %v", err, true)
		return
	}

	_, err = CheckFile("/")
	if err == nil {
		t.Errorf("CheckFile() error = %v, wantErr %v", err, true)
		return
	}

	_, err = CheckFile(".")
	if err == nil {
		t.Errorf("CheckFile() error = %v, wantErr %v", err, true)
		return
	}
}
