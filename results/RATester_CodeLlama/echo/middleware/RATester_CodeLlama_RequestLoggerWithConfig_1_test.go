package middleware

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestLoggerWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	config := RequestLoggerConfig{}

	// WHEN
	mw, err := config.ToMiddleware()

	// THEN
	require.NoError(t, err)
	require.NotNil(t, mw)
}
