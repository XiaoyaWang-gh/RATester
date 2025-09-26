package requestdecorator

import (
	"fmt"
	"testing"
	"time"
)

func TestLen_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := byTTL{
		&cnameResolv{
			TTL:    10 * time.Second,
			Record: "example.com",
		},
		&cnameResolv{
			TTL:    20 * time.Second,
			Record: "example.net",
		},
	}
	if len(a) != 2 {
		t.Errorf("expected len(a) == 2, got %d", len(a))
	}
}
