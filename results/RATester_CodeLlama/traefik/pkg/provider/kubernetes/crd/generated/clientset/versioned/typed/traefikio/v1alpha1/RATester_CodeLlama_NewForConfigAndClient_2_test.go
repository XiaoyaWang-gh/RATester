package v1alpha1

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	rest "k8s.io/client-go/rest"
)

func TestNewForConfigAndClient_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &rest.Config{}
	h := &http.Client{}

	// when
	client, err := NewForConfigAndClient(c, h)

	// then
	require.NoError(t, err)
	require.NotNil(t, client)
}
