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
		inclusions:    []glob.Glob{glob.MustCompile("*.go")},
		dirInclusions: []glob.Glob{glob.MustCompile("testdata")},
		isWindows:     false,
	}

	filter2 := &FilenameFilter{
		shouldInclude: func(filename string) bool { return true },
		inclusions:    []glob.Glob{glob.MustCompile("*.txt")},
		dirInclusions: []glob.Glob{glob.MustCompile("testdata")},
		isWindows:     false,
	}

	result := filter1.Append(filter2)

	if len(result.nested) != 2 {
		t.Errorf("Expected 2 nested filters, got %d", len(result.nested))
	}

	if result.nested[0] != filter1 {
		t.Errorf("Expected first nested filter to be filter1")
	}

	if result.nested[1] != filter2 {
		t.Errorf("Expected second nested filter to be filter2")
	}
}
