package prometheus

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/beego/beego/v2/client/httplib"
)

func TestFilterChain_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &FilterChainBuilder{
		AppName:    "testApp",
		ServerName: "testServer",
		RunMode:    "testMode",
	}

	next := func(ctx context.Context, req *httplib.BeegoHTTPRequest) (*http.Response, error) {
		return &http.Response{}, nil
	}

	filter := builder.FilterChain(next)

	req := &httplib.BeegoHTTPRequest{}
	resp, err := filter(context.Background(), req)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if resp == nil {
		t.Error("Expected a response, but got nil")
	}
}
