package middleware

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestContextTimeoutWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	config := ContextTimeoutConfig{
		Timeout: 10 * time.Second,
	}
	// WHEN
	mw, err := config.ToMiddleware()
	// THEN
	require.NoError(t, err)
	require.NotNil(t, mw)
}
