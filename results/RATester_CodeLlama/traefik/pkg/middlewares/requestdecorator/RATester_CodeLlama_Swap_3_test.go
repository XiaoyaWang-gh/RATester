package requestdecorator

import (
	"fmt"
	"testing"
	"time"
)

func TestSwap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := byTTL{
		&cnameResolv{
			TTL:    10 * time.Second,
			Record: "a.example.com",
		},
		&cnameResolv{
			TTL:    20 * time.Second,
			Record: "b.example.com",
		},
		&cnameResolv{
			TTL:    30 * time.Second,
			Record: "c.example.com",
		},
	}
	a.Swap(0, 1)
	if a[0].TTL != 20*time.Second || a[1].TTL != 10*time.Second {
		t.Errorf("Swap failed")
	}
}
