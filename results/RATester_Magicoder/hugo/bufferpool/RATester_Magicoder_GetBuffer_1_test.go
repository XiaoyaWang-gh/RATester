package bufferpool

import (
	"fmt"
	"testing"
)

func TestGetBuffer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := GetBuffer()
	if buf == nil {
		t.Error("Expected a non-nil buffer, got nil")
	}

	// Test that the buffer is empty
	if buf.Len() != 0 {
		t.Errorf("Expected buffer length to be 0, got %d", buf.Len())
	}

	// Test that the buffer is not full
	if buf.Cap() != 4096 {
		t.Errorf("Expected buffer capacity to be 4096, got %d", buf.Cap())
	}

	// Test that the buffer is available
	if buf.Available() != 4096 {
		t.Errorf("Expected buffer available to be 4096, got %d", buf.Available())
	}

	// Test that the buffer is available for writing
	if len(buf.AvailableBuffer()) != 4096 {
		t.Errorf("Expected buffer available buffer length to be 4096, got %d", len(buf.AvailableBuffer()))
	}

	// Test that the buffer is empty after reset
	buf.Reset()
	if buf.Len() != 0 {
		t.Errorf("Expected buffer length to be 0 after reset, got %d", buf.Len())
	}

	// Test that the buffer is not full after reset
	if buf.Cap() != 4096 {
		t.Errorf("Expected buffer capacity to be 4096 after reset, got %d", buf.Cap())
	}

	// Test that the buffer is available after reset
	if buf.Available() != 4096 {
		t.Errorf("Expected buffer available to be 4096 after reset, got %d", buf.Available())
	}

	// Test that the buffer is available for writing after reset
	if len(buf.AvailableBuffer()) != 4096 {
		t.Errorf("Expected buffer available buffer length to be 4096 after reset, got %d", len(buf.AvailableBuffer()))
	}
}
