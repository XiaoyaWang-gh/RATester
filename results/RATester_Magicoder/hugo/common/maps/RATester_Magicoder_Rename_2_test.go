package maps

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestRename_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	renamer := KeyRenamer{
		renames: []keyRename{
			{
				pattern: glob.MustCompile("oldKey*"),
				newKey:  "newKey",
			},
		},
	}

	m := map[string]any{
		"oldKey1":  "value1",
		"oldKey2":  "value2",
		"otherKey": "value3",
	}

	renamer.Rename(m)

	expected := map[string]any{
		"newKey1":  "value1",
		"newKey2":  "value2",
		"otherKey": "value3",
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Expected %v, got %v", expected, m)
	}
}
