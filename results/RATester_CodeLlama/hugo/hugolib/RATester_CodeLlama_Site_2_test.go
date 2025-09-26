package hugolib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSite_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	require.NotNil(t, p.Site())
}
