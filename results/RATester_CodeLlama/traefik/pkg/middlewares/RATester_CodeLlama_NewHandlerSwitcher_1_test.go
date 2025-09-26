package middlewares

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHandlerSwitcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	newHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	// when
	hs := NewHandlerSwitcher(newHandler)

	// then
	require.NotNil(t, hs)
	require.NotNil(t, hs.handler)
	require.Equal(t, newHandler, hs.handler.Get())
}
