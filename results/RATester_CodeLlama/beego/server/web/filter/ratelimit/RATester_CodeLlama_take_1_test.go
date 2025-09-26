package ratelimit

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestTake_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &limiter{}
	ctx := &context.Context{}
	amount := uint(1)
	if l.take(amount, ctx) {
		t.Errorf("l.take(amount, ctx) = %v, want %v", true, false)
	}
}
