package versioned

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	rest "k8s.io/client-go/rest"
)

func TestNewForConfigAndClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &rest.Config{}
	httpClient := &http.Client{}

	// when
	clientset, err := NewForConfigAndClient(c, httpClient)

	// then
	require.NoError(t, err)
	require.NotNil(t, clientset)
}
