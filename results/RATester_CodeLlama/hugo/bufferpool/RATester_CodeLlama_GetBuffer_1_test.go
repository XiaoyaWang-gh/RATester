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
		t.Errorf("GetBuffer() = nil, want not nil")
	}
	if buf.Len() != 0 {
		t.Errorf("GetBuffer().Len() = %d, want 0", buf.Len())
	}
	if buf.Cap() != 0 {
		t.Errorf("GetBuffer().Cap() = %d, want 0", buf.Cap())
	}
	if buf.Bytes() != nil {
		t.Errorf("GetBuffer().Bytes() = %v, want nil", buf.Bytes())
	}
	if buf.String() != "" {
		t.Errorf("GetBuffer().String() = %q, want \"\"", buf.String())
	}
}
