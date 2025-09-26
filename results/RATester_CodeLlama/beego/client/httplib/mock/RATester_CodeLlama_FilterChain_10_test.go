package mock

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/beego/beego/v2/client/httplib"
)

func TestFilterChain_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MockResponseFilter{}
	next := func(ctx context.Context, req *httplib.BeegoHTTPRequest) (*http.Response, error) {
		return &http.Response{}, nil
	}
	filter := m.FilterChain(next)
	ctx := context.Background()
	req := &httplib.BeegoHTTPRequest{}
	resp, err := filter(ctx, req)
	if resp == nil {
		t.Errorf("resp is nil")
	}
	if err != nil {
		t.Errorf("err is not nil")
	}
}
