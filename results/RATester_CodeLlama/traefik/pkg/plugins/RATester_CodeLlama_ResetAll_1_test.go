package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResetAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}

	err := c.ResetAll()
	require.Error(t, err)
	require.Equal(t, "goPath is empty", err.Error())
}
