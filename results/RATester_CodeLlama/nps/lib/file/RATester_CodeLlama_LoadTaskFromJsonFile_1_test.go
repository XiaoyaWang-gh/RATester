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

	// arrange
	var (
		err error
		s   = &JsonDb{
			RunPath:          "./",
			TaskFilePath:     "./task.json",
			ClientIncreaseId: 1,
			TaskIncreaseId:   1,
			HostIncreaseId:   1,
		}
	)
	// act
	s.LoadTaskFromJsonFile()
	// assert
	if err != nil {
		t.Errorf("LoadTaskFromJsonFile() error = %v", err)
	}
}
