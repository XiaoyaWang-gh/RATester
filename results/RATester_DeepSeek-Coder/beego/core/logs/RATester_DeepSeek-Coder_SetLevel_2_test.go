package logs

import (
	"fmt"
	"testing"
)

func TestSetLevel_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &BeeLogger{
		level: 0,
	}

	testLevel := 5
	logger.SetLevel(testLevel)

	if logger.level != testLevel {
		t.Errorf("Expected level to be %d, but got %d", testLevel, logger.level)
	}
}
