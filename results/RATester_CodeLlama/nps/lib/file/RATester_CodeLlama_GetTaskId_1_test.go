package file

import (
	"fmt"
	"testing"
)

func TestGetTaskId_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	//given
	var jsonDb JsonDb
	jsonDb.TaskIncreaseId = 100
	//when
	taskId := jsonDb.GetTaskId()
	//then
	if taskId != 101 {
		t.Errorf("GetTaskId() = %v, want %v", taskId, 101)
	}
}
