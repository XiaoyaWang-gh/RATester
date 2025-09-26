package limit

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/time/rate"
)

func TestWrite_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &Writer{
		w:       &bytes.Buffer{},
		limiter: rate.NewLimiter(rate.Inf, 10),
	}
	p := []byte("1234567890")
	n, err := w.Write(p)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(p) {
		t.Fatalf("n = %d, want %d", n, len(p))
	}
	if w.w.(*bytes.Buffer).Len() != len(p) {
		t.Fatalf("w.w.Len() = %d, want %d", w.w.(*bytes.Buffer).Len(), len(p))
	}
}
