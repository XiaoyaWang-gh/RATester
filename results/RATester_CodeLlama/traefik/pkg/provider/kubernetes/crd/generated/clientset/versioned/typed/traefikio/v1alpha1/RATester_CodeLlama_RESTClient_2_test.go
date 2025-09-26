package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRESTClient_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}

	// when
	result := c.RESTClient()

	// then
	require.NotNil(t, result)
}
