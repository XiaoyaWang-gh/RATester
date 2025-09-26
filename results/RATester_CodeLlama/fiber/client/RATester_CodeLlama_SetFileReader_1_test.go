package client

import (
	"fmt"
	"io"
	"testing"
)

func TestSetFileReader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var r io.ReadCloser
	var f *File
	// act
	SetFileReader(r)(f)
	// assert
	if f.reader != r {
		t.Errorf("SetFileReader() failed")
	}
}
