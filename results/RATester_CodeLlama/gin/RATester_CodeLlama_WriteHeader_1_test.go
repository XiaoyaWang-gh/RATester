package gin

import (
	"fmt"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{
		status: 200,
	}
	w.WriteHeader(200)
	if w.status != 200 {
		t.Errorf("w.status = %d; want 200", w.status)
	}
	w.WriteHeader(404)
	if w.status != 404 {
		t.Errorf("w.status = %d; want 404", w.status)
	}
	w.WriteHeader(200)
	if w.status != 404 {
		t.Errorf("w.status = %d; want 404", w.status)
	}
}
