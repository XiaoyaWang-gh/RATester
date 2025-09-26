package crd

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestNamespaceOrFallback_1(t *testing.T) {
	testCases := []struct {
		desc     string
		lb       traefikv1alpha1.LoadBalancerSpec
		fallback string
		expected string
	}{
		{
			desc:     "should return the namespace when it is set",
			lb:       traefikv1alpha1.LoadBalancerSpec{Namespace: "namespace"},
			fallback: "fallback",
			expected: "namespace",
		},
		{
			desc:     "should return the fallback when the namespace is not set",
			lb:       traefikv1alpha1.LoadBalancerSpec{},
			fallback: "fallback",
			expected: "fallback",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			actual := namespaceOrFallback(test.lb, test.fallback)
			assert.Equal(t, test.expected, actual)
		})
	}
}
