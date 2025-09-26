package httplib

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithTLSClientConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	var config *tls.Config
	// WHEN
	option := WithTLSClientConfig(config)
	// THEN
	require.NotNil(t, option)
}
