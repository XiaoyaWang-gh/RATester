package hugo

import (
	"fmt"
	"testing"
)

func TestReleaseVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Version{Major: 1, Minor: 2, PatchLevel: 3, Suffix: "test"}
	result := v.ReleaseVersion()
	if result.Suffix != "" {
		t.Errorf("Expected Suffix to be empty, but got %s", result.Suffix)
	}
}
