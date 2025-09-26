package hugio

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestNewReadSeekerNoOpCloserFromBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	content := []byte("test content")
	rs := NewReadSeekerNoOpCloserFromBytes(content)

	// Test Read
	buf := make([]byte, len(content))
	n, err := rs.Read(buf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(content) {
		t.Errorf("Expected to read %d bytes, but read %d", len(content), n)
	}
	if !bytes.Equal(buf, content) {
		t.Errorf("Expected to read %v, but read %v", content, buf)
	}

	// Test Seek
	offset, err := rs.Seek(0, io.SeekStart)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if offset != 0 {
		t.Errorf("Expected to seek to 0, but seeked to %d", offset)
	}

	// Test Close
	err = rs.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
