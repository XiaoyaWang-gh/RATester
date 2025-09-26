package file

import (
	"fmt"
	"testing"
)

func TestLoadTaskFromJsonFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		TaskFilePath: "testdata/tasks.json",
	}

	db.LoadTaskFromJsonFile()

	// Add assertions here to verify the correctness of the method
}
