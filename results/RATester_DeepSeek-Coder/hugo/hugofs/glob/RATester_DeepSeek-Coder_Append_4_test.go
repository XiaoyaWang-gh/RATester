package glob

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestAppend_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filter1 := &FilenameFilter{
		shouldInclude: func(filename string) bool { return true },
		inclusions:    []glob.Glob{glob.MustCompile("*")},
		dirInclusions: []glob.Glob{glob.MustCompile("*")},
		exclusions:    []glob.Glob{glob.MustCompile("*")},
		isWindows:     true,
		nested:        []*FilenameFilter{},
	}

	filter2 := &FilenameFilter{
		shouldInclude: func(filename string) bool { return false },
		inclusions:    []glob.Glob{glob.MustCompile("*")},
		dirInclusions: []glob.Glob{glob.MustCompile("*")},
		exclusions:    []glob.Glob{glob.MustCompile("*")},
		isWindows:     false,
		nested:        []*FilenameFilter{},
	}

	result := filter1.Append(filter2)

	if result == nil {
		t.Errorf("Expected result to not be nil")
	}

	if len(result.nested) != 1 {
		t.Errorf("Expected result.nested to have length 1, got %d", len(result.nested))
	}

	if result.nested[0] != filter2 {
		t.Errorf("Expected result.nested[0] to be filter2")
	}
}
