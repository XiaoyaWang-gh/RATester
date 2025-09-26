package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rp *Provider
	if rp.SessionAll(context.Background()) != 0 {
		t.Errorf("SessionAll() = %v, want %v", rp.SessionAll(context.Background()), 0)
	}
}
