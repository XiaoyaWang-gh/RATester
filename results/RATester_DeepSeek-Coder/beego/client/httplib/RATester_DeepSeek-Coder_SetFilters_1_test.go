package httplib

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSetFilters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		url: "http://example.com",
	}

	filter1 := func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			// do something before
			resp, err := next(ctx, req)
			// do something after
			return resp, err
		}
	}

	filter2 := func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			// do something before
			resp, err := next(ctx, req)
			// do something after
			return resp, err
		}
	}

	b.SetFilters(filter1, filter2)

	// assert something about b.setting.FilterChains
}
