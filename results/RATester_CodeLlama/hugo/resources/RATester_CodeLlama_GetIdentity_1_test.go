package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/identity"
)

func TestGetIdentity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var r *resourceAdapter
	// Act
	result := r.GetIdentity()
	// Assert
	assert.Equal(t, identity.FirstIdentity(r.target), result)
}
