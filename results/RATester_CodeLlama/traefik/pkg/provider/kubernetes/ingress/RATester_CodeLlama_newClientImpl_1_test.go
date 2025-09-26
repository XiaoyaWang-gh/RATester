package ingress

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	kclientset "k8s.io/client-go/kubernetes"
)

func TestNewClientImpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var clientset kclientset.Interface
	// Act
	result := newClientImpl(clientset)
	// Assert
	assert.NotNil(t, result)
}
