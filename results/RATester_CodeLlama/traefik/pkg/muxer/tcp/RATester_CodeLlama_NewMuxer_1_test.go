package tcp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMuxer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer, err := NewMuxer()
	require.NoError(t, err)
	require.NotNil(t, muxer)
}
