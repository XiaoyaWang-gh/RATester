package gateway

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	gatev1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFilterReferenceGrantsFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var referenceGrants []*gatev1beta1.ReferenceGrant
	var group string
	var kind string
	var namespace string
	// Act
	var matchingReferenceGrants []*gatev1beta1.ReferenceGrant
	matchingReferenceGrants = filterReferenceGrantsFrom(referenceGrants, group, kind, namespace)
	// Assert
	assert.Equal(t, matchingReferenceGrants, []*gatev1beta1.ReferenceGrant{})
}
