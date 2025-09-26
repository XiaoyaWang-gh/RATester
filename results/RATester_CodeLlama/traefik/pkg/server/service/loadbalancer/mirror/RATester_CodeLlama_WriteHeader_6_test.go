package mirror

import (
	"fmt"
	"testing"
)

func TestWriteHeader_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	b.WriteHeader(1)
	if b.Header().Get("Content-Length") != "" {
		t.Errorf("WriteHeader() should not set Content-Length header")
	}
}
