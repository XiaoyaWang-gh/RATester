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

	p := &Provider{
		client: &ssdb.Client{},
	}
	if p.SessionAll(context.Background()) != 0 {
		t.Error("SessionAll() failed")
	}
}
