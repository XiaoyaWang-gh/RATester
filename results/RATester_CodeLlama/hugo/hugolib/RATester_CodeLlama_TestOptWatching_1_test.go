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

	c := &IntegrationTestConfig{}
	TestOptWatching()(c)
	if c.Watching != true {
		t.Errorf("TestOptWatching() failed")
	}
}
