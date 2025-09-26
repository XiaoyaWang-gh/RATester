package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_41(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	deps := &deps.Deps{}

	// Act
	_, err := New(deps)

	// Assert
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
}
