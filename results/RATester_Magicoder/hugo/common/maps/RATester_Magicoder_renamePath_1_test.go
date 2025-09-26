package maps

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestrenamePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	renamer := KeyRenamer{
		renames: []keyRename{
			{
				pattern: glob.MustCompile("foo.*"),
				newKey:  "bar",
			},
		},
	}

	input := map[string]any{
		"foo.bar": "baz",
		"foo.baz": "qux",
	}

	expected := map[string]any{
		"bar":     "baz",
		"foo.baz": "qux",
	}

	renamer.renamePath("", input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}
