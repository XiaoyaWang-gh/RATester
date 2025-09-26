package hugio

import (
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HasBytesWriter{
		Patterns: []*HasBytesPattern{
			{
				Match:   false,
				Pattern: []byte("foo"),
			},
			{
				Match:   false,
				Pattern: []byte("bar"),
			},
		},
	}

	p := []byte("foobar")
	n, err := h.Write(p)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if n != len(p) {
		t.Errorf("expected %d, got %d", len(p), n)
	}

	if !h.done {
		t.Errorf("expected done to be true")
	}

	if !h.Patterns[0].Match {
		t.Errorf("expected Patterns[0].Match to be true")
	}

	if !h.Patterns[1].Match {
		t.Errorf("expected Patterns[1].Match to be true")
	}
}
