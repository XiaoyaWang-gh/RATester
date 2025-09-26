package maps

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestgetNewKey_1(t *testing.T) {
	renamer := KeyRenamer{
		renames: []keyRename{
			{
				pattern: glob.MustCompile("oldKey"),
				newKey:  "newKey",
			},
		},
	}

	tests := []struct {
		name     string
		keyPath  string
		expected string
	}{
		{
			name:     "Matching key",
			keyPath:  "oldKey",
			expected: "newKey",
		},
		{
			name:     "Non-matching key",
			keyPath:  "otherKey",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := renamer.getNewKey(test.keyPath)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
