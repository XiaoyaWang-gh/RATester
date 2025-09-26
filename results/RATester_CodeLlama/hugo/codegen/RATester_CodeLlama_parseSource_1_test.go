package codegen

import (
	"fmt"
	"testing"
)

func TestParseSource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Inspector{
		ProjectRootDir: "hugo",
	}
	c.parseSource()
	if len(c.methodWeight) == 0 {
		t.Errorf("Expected methodWeight to be populated")
	}
}
