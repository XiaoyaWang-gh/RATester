package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOpenFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &OpenFilesFs{
		openFiles: map[string]int{
			"file1": 1,
			"file2": 2,
		},
	}

	expected := map[string]int{
		"file1": 1,
		"file2": 2,
	}

	result := fs.OpenFiles()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
