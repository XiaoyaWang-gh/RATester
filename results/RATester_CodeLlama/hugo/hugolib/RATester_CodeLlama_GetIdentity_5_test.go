package hugolib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetIdentity_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	require.Equal(t, p, p.GetIdentity())
}
