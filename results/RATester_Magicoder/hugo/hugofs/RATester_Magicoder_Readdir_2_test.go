package hugofs

import (
	"fmt"
	"testing"
)

func TestReaddir_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := &rootMappingDir{
		name: "testDir",
	}

	_, err := dir.Readdir(10)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Error() != "not supported: use ReadDir" {
		t.Errorf("Expected error message 'not supported: use ReadDir', but got '%s'", err.Error())
	}
}
