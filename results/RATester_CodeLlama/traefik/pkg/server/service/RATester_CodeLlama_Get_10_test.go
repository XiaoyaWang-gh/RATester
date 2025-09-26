package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RoundTripperManager{
		roundTrippers: map[string]http.RoundTripper{
			"default@internal": &http.Transport{},
		},
	}

	rt, err := r.Get("")
	require.NoError(t, err)
	require.NotNil(t, rt)

	rt, err = r.Get("default@internal")
	require.NoError(t, err)
	require.NotNil(t, rt)

	rt, err = r.Get("default@external")
	require.Error(t, err)
	require.Nil(t, rt)
}
