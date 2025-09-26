package hugolib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	require.Nil(t, p.Err())
}
