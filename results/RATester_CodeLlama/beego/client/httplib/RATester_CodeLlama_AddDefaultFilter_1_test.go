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
			return nil, nil
		}
	}
	AddDefaultFilter(fc)
	if len(defaultSetting.FilterChains) != 1 {
		t.Errorf("AddDefaultFilter failed")
	}
}
