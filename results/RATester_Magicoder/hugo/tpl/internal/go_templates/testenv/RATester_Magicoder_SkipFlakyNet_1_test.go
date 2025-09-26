package testenv

import (
	"fmt"
	"os"
	"testing"
)

func TestSkipFlakyNet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("GO_BUILDER_FLAKY_NET", "true")
	SkipFlakyNet(t)
	if !t.Skipped() {
		t.Errorf("Expected test to be skipped")
	}
}
