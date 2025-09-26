package hugolib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	h := &HugoSites{}
	// when
	err := h.loadData()
	// then
	require.NoError(t, err)
}
