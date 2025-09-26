package modules

import (
	"fmt"
	"testing"
)

func TestLoadModules_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &collector{}
	err := c.loadModules()
	if err != nil {
		t.Errorf("loadModules() error = %v", err)
		return
	}
}
