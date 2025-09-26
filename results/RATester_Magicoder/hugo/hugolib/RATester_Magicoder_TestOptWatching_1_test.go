package hugolib

import (
	"fmt"
	"testing"
)

func TestTestOptWatching_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &IntegrationTestConfig{}
	TestOptWatching()(config)

	if !config.Watching {
		t.Errorf("Expected Watching to be true, but it was false")
	}
}
