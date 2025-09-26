package redirect

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewRequestRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	conf := dynamic.RequestRedirect{
		Scheme:     nil,
		Hostname:   nil,
		Port:       nil,
		Path:       nil,
		PathPrefix: nil,
		StatusCode: 0,
	}
	name := "test"

	// WHEN
	handler, err := NewRequestRedirect(ctx, next, conf, name)

	// THEN
	require.NoError(t, err)
	require.NotNil(t, handler)
}
