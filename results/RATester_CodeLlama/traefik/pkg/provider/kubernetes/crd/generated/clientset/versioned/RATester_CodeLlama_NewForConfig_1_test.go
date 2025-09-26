package versioned

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	rest "k8s.io/client-go/rest"
)

func TestNewForConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &rest.Config{}

	// when
	clientset, err := NewForConfig(c)

	// then
	require.NoError(t, err)
	require.NotNil(t, clientset)
}
