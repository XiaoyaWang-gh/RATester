package filecache

import (
	"fmt"
	"testing"
)

func TestAsHTTPCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{}
	hc := c.AsHTTPCache()
	if hc == nil {
		t.Fatal("AsHTTPCache returned nil")
	}
}
