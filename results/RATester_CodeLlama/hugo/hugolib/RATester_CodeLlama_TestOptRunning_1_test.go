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

	c := &IntegrationTestConfig{}
	TestOptRunning()(c)
	if !c.Running {
		t.Errorf("TestOptRunning() failed")
	}
}
