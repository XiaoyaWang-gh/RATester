package resources

import (
	"fmt"
	"testing"
)

func TestInitTransform_1(t *testing.T) {
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
	r.initTransform(publish, setContent)

	// Assert
	// ...
}
