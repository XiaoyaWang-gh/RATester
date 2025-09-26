package capture

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWrite_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{}
	b := []byte("Hello, World!")
	size, err := crw.Write(b)
	if err != nil {
		t.Errorf("Write() error = %v", err)
		return
	}
	if size != len(b) {
		t.Errorf("Write() size = %v, want %v", size, len(b))
	}
	if crw.size != int64(size) {
		t.Errorf("Write() crw.size = %v, want %v", crw.size, size)
	}
	if crw.status != http.StatusOK {
		t.Errorf("Write() crw.status = %v, want %v", crw.status, http.StatusOK)
	}
}
