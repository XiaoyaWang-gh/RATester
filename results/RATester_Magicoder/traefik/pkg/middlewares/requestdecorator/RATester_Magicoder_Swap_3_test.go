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
		&cnameResolv{TTL: 10 * time.Second, Record: "record1"},
		&cnameResolv{TTL: 20 * time.Second, Record: "record2"},
	}

	a.Swap(0, 1)

	if a[0].TTL != 20*time.Second || a[0].Record != "record2" {
		t.Errorf("Expected TTL: 20s, Record: record2. Got TTL: %s, Record: %s", a[0].TTL, a[0].Record)
	}

	if a[1].TTL != 10*time.Second || a[1].Record != "record1" {
		t.Errorf("Expected TTL: 10s, Record: record1. Got TTL: %s, Record: %s", a[1].TTL, a[1].Record)
	}
}
