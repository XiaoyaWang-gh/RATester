package files

import (
	"fmt"
	"sort"
	"testing"
)

func Testinit_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ComponentFolders = []string{"folder1", "folder2", "folder3"}
	sort.Strings(ComponentFolders)
	for _, f := range ComponentFolders {
		componentFoldersSet[f] = true
	}

	// Test if the ComponentFolders is sorted
	for i := 0; i < len(ComponentFolders)-1; i++ {
		if ComponentFolders[i] > ComponentFolders[i+1] {
			t.Errorf("ComponentFolders is not sorted")
		}
	}

	// Test if all elements in ComponentFolders are in componentFoldersSet
	for _, f := range ComponentFolders {
		if _, ok := componentFoldersSet[f]; !ok {
			t.Errorf("Element %s in ComponentFolders is not in componentFoldersSet", f)
		}
	}
}
