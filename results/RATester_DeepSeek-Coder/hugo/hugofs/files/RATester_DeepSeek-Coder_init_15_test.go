package files

import (
	"fmt"
	"sort"
	"testing"
)

func TestInit_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sort.Strings(ComponentFolders)
	for _, f := range ComponentFolders {
		if _, ok := componentFoldersSet[f]; !ok {
			t.Errorf("Expected %s to be in componentFoldersSet", f)
		}
	}
}
