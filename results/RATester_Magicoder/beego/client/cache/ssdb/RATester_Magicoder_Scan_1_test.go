package ssdb

import (
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestScan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &Cache{
		conn:     &ssdb.Client{},
		conninfo: []string{},
	}

	keyStart := "keyStart"
	keyEnd := "keyEnd"
	limit := 10

	resp, err := rc.Scan(keyStart, keyEnd, limit)
	if err != nil {
		t.Errorf("Scan failed, error: %v", err)
	}

	if len(resp) != limit {
		t.Errorf("Expected %d results, got %d", limit, len(resp))
	}
}
