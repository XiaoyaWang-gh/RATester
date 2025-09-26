package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestSessionGC_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{
		client:      &ssdb.Client{},
		Host:        "127.0.0.1",
		Port:        8888,
		maxLifetime: 1000,
	}
	ctx := context.Background()
	p.SessionGC(ctx)
}
