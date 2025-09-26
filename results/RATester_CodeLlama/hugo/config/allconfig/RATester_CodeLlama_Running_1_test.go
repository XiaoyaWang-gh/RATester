package allconfig

import (
	"fmt"
	"testing"
)

func TestRunning_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.Running() {
		t.Error("Running() should be false")
	}
}
