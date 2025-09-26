package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildHTTP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rootCtx := context.Background()
	serviceName := "test"
	m := &Manager{}

	// when
	handler, err := m.BuildHTTP(rootCtx, serviceName)

	// then
	require.NoError(t, err)
	require.NotNil(t, handler)
}
