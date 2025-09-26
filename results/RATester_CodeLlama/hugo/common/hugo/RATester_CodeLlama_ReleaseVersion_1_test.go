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

	v := Version{
		Major:      1,
		Minor:      2,
		PatchLevel: 3,
		Suffix:     "dev",
	}

	v.ReleaseVersion()

	if v.Suffix != "" {
		t.Errorf("Expected empty suffix, got %q", v.Suffix)
	}
}
