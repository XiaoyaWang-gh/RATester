package v1

import (
	"fmt"
	"testing"
)

func TestGetBaseConfig_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &VisitorBaseConfig{
		Name: "test",
	}
	if c.GetBaseConfig() == nil {
		t.Fatal("GetBaseConfig() failed")
	}
}
