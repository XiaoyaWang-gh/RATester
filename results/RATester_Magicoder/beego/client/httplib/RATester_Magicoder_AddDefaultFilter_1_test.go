package httplib

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestAddDefaultFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			// Your filter logic here
			return next(ctx, req)
		}
	}

	AddDefaultFilter(fc)

	// Assert that the filter has been added to the default setting
	if len(defaultSetting.FilterChains) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(defaultSetting.FilterChains))
	}
}
