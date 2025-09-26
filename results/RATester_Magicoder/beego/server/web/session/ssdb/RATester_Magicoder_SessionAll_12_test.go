package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestSessionAll_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	p := &Provider{
		client:      &ssdb.Client{},
		Host:        "localhost",
		Port:        8888,
		maxLifetime: 3600,
	}
	result := p.SessionAll(ctx)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
