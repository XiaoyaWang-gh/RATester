package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRESTClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	fake := &FakeTraefikV1alpha1{}
	// when
	restClient := fake.RESTClient()
	// then
	require.NotNil(t, restClient)
}
