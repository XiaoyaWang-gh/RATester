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

	b := &BeegoHTTPRequest{}
	b.AddFilters(func(next Filter) Filter {
		return func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
			return nil, nil
		}
	})
	if len(b.setting.FilterChains) != 1 {
		t.Errorf("TestAddFilters failed")
	}
}
