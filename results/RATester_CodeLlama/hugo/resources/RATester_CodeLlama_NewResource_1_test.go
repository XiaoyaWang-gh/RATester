package resources

import (
	"fmt"
	"testing"
)

func TestNewResource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var r *Spec
	var rd ResourceSourceDescriptor

	// Act
	_, err := r.NewResource(rd)

	// Assert
	if err != nil {
		t.Errorf("NewResource() error = %v", err)
		return
	}
}
