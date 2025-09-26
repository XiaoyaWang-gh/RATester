package capture

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	ctx := context.Background()
	ctx = context.WithValue(ctx, capturedData, &Capture{})

	// WHEN
	capture, err := FromContext(ctx)

	// THEN
	require.NoError(t, err)
	require.NotNil(t, capture)
}
