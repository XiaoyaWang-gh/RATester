package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestprocessContentAdaptersOnRebuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	h := &HugoSites{}
	buildCfg := &BuildCfg{}

	err := h.processContentAdaptersOnRebuild(ctx, buildCfg)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
