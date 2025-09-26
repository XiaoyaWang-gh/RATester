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

	db := &JsonDb{
		TaskIncreaseId: 0,
	}

	expected := int32(1)
	actual := db.GetTaskId()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	expected = int32(2)
	actual = db.GetTaskId()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
