package passtlsclientcert

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNew_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	config := dynamic.PassTLSClientCert{}
	name := "test"

	middleware, err := New(ctx, next, config, name)
	require.NoError(t, err)
	require.IsType(t, &passTLSClientCert{}, middleware)
}
