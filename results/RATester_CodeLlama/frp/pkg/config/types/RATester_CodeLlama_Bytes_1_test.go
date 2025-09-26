package types

import (
	"fmt"
	"testing"
)

func TestBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &BandwidthQuantity{
		s: "100MB",
		i: 100 * 1024 * 1024,
	}
	if q.Bytes() != 100*1024*1024 {
		t.Fatal("q.Bytes() != 100*1024*1024")
	}
}
