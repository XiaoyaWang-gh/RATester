package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestSessionExist_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{
		client: &ssdb.Client{},
	}
	ctx := context.Background()
	sid := "123"
	value, err := p.SessionExist(ctx, sid)
	if err != nil {
		t.Error(err)
	}
	if value != true {
		t.Error("SessionExist failed")
	}
}
