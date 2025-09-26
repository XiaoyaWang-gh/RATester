package hugofs

import (
	"fmt"
	"testing"
)

func TestReverseLookup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var (
		fs = new(RootMappingFs)
		// CONTEXT_0
		filename = "foo.md"
	)
	// CONTEXT_1
	// CONTEXT_2
	// CONTEXT_3
	// ACT
	_, err := fs.ReverseLookup(filename)
	// ASSERT
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
