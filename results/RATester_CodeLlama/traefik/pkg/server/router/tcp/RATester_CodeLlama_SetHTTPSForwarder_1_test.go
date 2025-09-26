package tcp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestSetHTTPSForwarder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Router{}
	handler := &tcp.TLSHandler{}

	// when
	r.SetHTTPSForwarder(handler)

	// then
	require.NotNil(t, r.httpsForwarder)
	require.Equal(t, handler, r.httpsForwarder)
}
