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

	b := &BeegoHTTPRequest{}

	// Create a mock filter
	mockFilter := func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			// Add some logic here
			return next(ctx, req)
		}
	}

	// Set the filter
	b.SetFilters(mockFilter)

	// Add more filters
	b.SetFilters(mockFilter, mockFilter)

	// Check if the filters are set correctly
	if len(b.setting.FilterChains) != 2 {
		t.Errorf("Expected 2 filters, got %d", len(b.setting.FilterChains))
	}
}
