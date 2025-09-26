package deps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Deps{}
	err := d.Init()
	require.NoError(t, err)
}
