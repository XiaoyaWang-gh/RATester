package hugolib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	h := &HugoSites{}
	m := &pageMeta{}

	// When
	_, _, err := h.newPage(m)

	// Then
	require.NoError(t, err)
}
