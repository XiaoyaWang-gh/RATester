package ratelimit

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestdefaultSessionKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context.Context{}
	result := defaultSessionKey(ctx)
	if result != "BEEGO_ALL" {
		t.Errorf("Expected 'BEEGO_ALL', but got '%s'", result)
	}
}
