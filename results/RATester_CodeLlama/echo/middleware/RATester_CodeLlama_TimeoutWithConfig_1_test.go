package middleware

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeoutWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	config := TimeoutConfig{
		Timeout: 100 * time.Millisecond,
	}
	// WHEN
	mw, err := config.ToMiddleware()
	// THEN
	require.NoError(t, err)
	require.NotNil(t, mw)
}
