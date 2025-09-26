package file

import (
	"fmt"
	"testing"
)

func TestLoadHostFromJsonFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		HostFilePath: "testdata/hosts.json",
	}
	db.LoadHostFromJsonFile()

	// Add your assertions here
}
