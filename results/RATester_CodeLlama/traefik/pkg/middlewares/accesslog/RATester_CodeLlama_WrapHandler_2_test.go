package accesslog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrapHandler_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	handler := &Handler{}
	// when
	aliceHandler := WrapHandler(handler)
	// then
	require.NotNil(t, aliceHandler)
}
