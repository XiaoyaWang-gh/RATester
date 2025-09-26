package ratelimit

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestDefaultSessionKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context.Context{}
	if got := defaultSessionKey(ctx); got != "BEEGO_ALL" {
		t.Errorf("defaultSessionKey() = %v, want %v", got, "BEEGO_ALL")
	}
}
