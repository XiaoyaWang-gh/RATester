package media

import (
	"fmt"
	"testing"
)

func TestString_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{
		Type: "application/rss+xml",
	}
	if m.String() != "application/rss+xml" {
		t.Errorf("m.String() = %v, want %v", m.String(), "application/rss+xml")
	}
}
