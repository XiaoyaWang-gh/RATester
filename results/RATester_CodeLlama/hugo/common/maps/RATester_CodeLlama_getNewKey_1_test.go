package maps

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestGetNewKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	keyPath := "a/b/c"
	r := KeyRenamer{
		renames: []keyRename{
			{
				pattern: glob.MustCompile("a/b/c"),
				newKey:  "d/e/f",
			},
		},
	}

	// Act
	newKey := r.getNewKey(keyPath)

	// Assert
	if newKey != "d/e/f" {
		t.Errorf("Expected newKey to be %q but was %q", "d/e/f", newKey)
	}
}
