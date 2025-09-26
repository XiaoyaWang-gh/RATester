package limit

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Reader{
		r:       bytes.NewReader([]byte("hello")),
		limiter: rate.NewLimiter(rate.Every(time.Second), 1),
	}
	p := make([]byte, 10)
	n, err := r.Read(p)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatalf("read %d bytes, want 5", n)
	}
	if string(p[:n]) != "hello" {
		t.Fatalf("read %q, want %q", string(p[:n]), "hello")
	}
}
