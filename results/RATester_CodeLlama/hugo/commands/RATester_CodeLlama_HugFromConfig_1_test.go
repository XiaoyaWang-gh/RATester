package commands

import (
	"fmt"
	"testing"
)

func TestHugFromConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	conf := &commonConfig{}
	r := &rootCommand{}
	// Act
	_, err := r.HugFromConfig(conf)
	// Assert
	if err != nil {
		t.Errorf("HugFromConfig() error = %v", err)
		return
	}
}
