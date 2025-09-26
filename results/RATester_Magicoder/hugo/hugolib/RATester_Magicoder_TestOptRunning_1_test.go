package hugolib

import (
	"fmt"
	"testing"
)

func TestTestOptRunning_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &IntegrationTestConfig{}
	TestOptRunning()(config)

	if !config.Running {
		t.Error("Expected Running to be true, but it was false")
	}
}
