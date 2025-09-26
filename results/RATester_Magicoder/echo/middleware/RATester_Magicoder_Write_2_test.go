package middleware

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &bodyDumpResponseWriter{
		Writer: &bytes.Buffer{},
	}
	testData := []byte("test data")
	n, err := w.Write(testData)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, but wrote %d", len(testData), n)
	}
	writtenData := w.Writer.(*bytes.Buffer).Bytes()
	if !bytes.Equal(testData, writtenData) {
		t.Errorf("Expected to write %v, but wrote %v", testData, writtenData)
	}
}
