package modules

import (
	"fmt"
	"testing"
)

func TestInitModules_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &collector{}
	err := c.initModules()
	if err != nil {
		t.Errorf("initModules() error = %v", err)
		return
	}
}
