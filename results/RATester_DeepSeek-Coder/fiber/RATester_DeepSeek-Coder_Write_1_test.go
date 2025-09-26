package fiber

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tc := &testConn{}
	testData := []byte("test data")

	n, err := tc.Write(testData)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	readData := make([]byte, len(testData))
	n, err = tc.Read(readData)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(testData) {
		t.Errorf("Expected to read %d bytes, read %d", len(testData), n)
	}

	if !bytes.Equal(testData, readData) {
		t.Errorf("Expected to read data %v, read %v", testData, readData)
	}
}
