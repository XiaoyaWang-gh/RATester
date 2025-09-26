package deps

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGetFilenamesWithPostPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{
		filenamesWithPostPrefix: map[string]bool{
			"file1": true,
			"file2": true,
			"file3": true,
		},
	}

	filenames := b.GetFilenamesWithPostPrefix()

	sort.Strings(filenames)

	expected := []string{"file1", "file2", "file3"}

	if !reflect.DeepEqual(filenames, expected) {
		t.Errorf("Expected %v, got %v", expected, filenames)
	}
}
