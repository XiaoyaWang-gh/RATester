package hugolib

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestTestOptDebug_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IntegrationTestConfig{}
	TestOptDebug()(c)
	if c.LogLevel != logg.LevelDebug {
		t.Errorf("TestOptDebug() failed")
	}
}
