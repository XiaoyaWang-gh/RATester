package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRead_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{
		client:      nil,
		Host:        "127.0.0.1",
		Port:        8888,
		maxLifetime: 60,
	}
	ctx := context.Background()
	sid := "123456"
	_, err := p.SessionRead(ctx, sid)
	if err != nil {
		t.Error(err)
	}
}
