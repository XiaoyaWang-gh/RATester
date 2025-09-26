package allconfig

import (
	"fmt"
	"testing"
)

func TestEnvironment_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.Environment() != "" {
		t.Errorf("Expected empty string, got %s", c.Environment())
	}
}
