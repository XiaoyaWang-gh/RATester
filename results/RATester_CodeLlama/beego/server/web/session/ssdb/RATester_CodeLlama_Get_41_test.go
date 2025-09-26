package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_41(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionStore{
		sid: "123456",
	}
	ctx := context.Background()
	key := "key"
	value := "value"
	s.values[key] = value
	if s.Get(ctx, key) != value {
		t.Errorf("Get() = %v, want %v", s.Get(ctx, key), value)
	}
}
