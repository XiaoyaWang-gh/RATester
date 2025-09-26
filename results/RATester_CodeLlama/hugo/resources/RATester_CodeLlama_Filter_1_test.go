package resources

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var i *imageResource
	var filters []any

	// when
	result, err := i.Filter(filters...)

	// then
	require.NoError(t, err)
	require.NotNil(t, result)
}
