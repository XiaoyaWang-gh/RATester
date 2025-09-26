package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestGetEntryPoints_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rw := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	h := Handler{
		staticConfig: static.Configuration{
			EntryPoints: map[string]*static.EntryPoint{
				"entrypoint1": {
					Address: "address1",
				},
				"entrypoint2": {
					Address: "address2",
				},
			},
		},
	}

	// when
	h.getEntryPoints(rw, request)

	// then
	require.Equal(t, http.StatusOK, rw.Code)
	require.Equal(t, "application/json", rw.Header().Get("Content-Type"))
	require.Equal(t, `{"entrypoints":[{"address":"address1","name":"entrypoint1"},{"address":"address2","name":"entrypoint2"}]}`, rw.Body.String())
}
