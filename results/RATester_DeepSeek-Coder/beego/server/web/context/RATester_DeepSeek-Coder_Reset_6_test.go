package context

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReset_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testWriter := &bytes.Buffer{}
	nrw := nopResetWriter{}

	nrw.Reset(testWriter)

	testString := "Hello, World!"
	n, err := nrw.Write([]byte(testString))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(testString) {
		t.Errorf("Expected to write %d bytes, but wrote %d", len(testString), n)
	}

	if testWriter.String() != testString {
		t.Errorf("Expected to write '%s', but wrote '%s'", testString, testWriter.String())
	}
}
