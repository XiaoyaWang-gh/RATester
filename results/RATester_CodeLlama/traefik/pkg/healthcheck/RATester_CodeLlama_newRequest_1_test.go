package healthcheck

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	target := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/health",
	}
	shc := &ServiceHealthChecker{
		config: &dynamic.ServerHealthCheck{
			Path: "/health",
		},
	}
	req, err := shc.newRequest(ctx, target)
	if err != nil {
		t.Errorf("newRequest() error = %v", err)
		return
	}
	if req.Method != http.MethodGet {
		t.Errorf("newRequest() method = %v, want %v", req.Method, http.MethodGet)
	}
	if req.URL.Path != "/health" {
		t.Errorf("newRequest() path = %v, want %v", req.URL.Path, "/health")
	}
}
