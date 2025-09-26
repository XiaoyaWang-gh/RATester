package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var mp *Provider
	if mp.SessionAll(context.Background()) != 0 {
		t.Errorf("SessionAll() = %v, want %v", mp.SessionAll(context.Background()), 0)
	}
}
