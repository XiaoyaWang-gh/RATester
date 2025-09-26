package resources

import (
	"fmt"
	"testing"
)

func TestGetOrTransform_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var r *resourceAdapter
	var publish bool
	var setContent bool

	// Act
	err := r.getOrTransform(publish, setContent)

	// Assert
	if err != nil {
		t.Errorf("getOrTransform() error = %v, wantErr %v", err, nil)
		return
	}
}
