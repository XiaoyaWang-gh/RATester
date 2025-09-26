package httplib

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestAddFilters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		setting: BeegoHTTPSettings{
			FilterChains: make([]FilterChain, 0),
		},
	}

	f1 := func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			// do something
			return next(ctx, req)
		}
	}

	f2 := func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			// do something
			return next(ctx, req)
		}
	}

	b.AddFilters(f1, f2)

	if len(b.setting.FilterChains) != 2 {
		t.Errorf("Expected 2 filters, got %d", len(b.setting.FilterChains))
	}
}
