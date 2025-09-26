package grpcweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNew_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	config := dynamic.GrpcWeb{
		AllowOrigins: []string{"*"},
	}
	name := "test"

	handler := New(ctx, next, config, name)

	assert.NotNil(t, handler)
}
