package redirect

import (
	"fmt"
	"testing"
)

func TestConfigDefault_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			StatusCode: 301,
		},
	}
	// Act
	result := configDefault(config...)
	// Assert
	if result.StatusCode != 301 {
		t.Errorf("Expected %d, got %d", 301, result.StatusCode)
	}
}
