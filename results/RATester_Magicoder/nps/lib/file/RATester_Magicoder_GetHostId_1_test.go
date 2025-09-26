package file

import (
	"fmt"
	"testing"
)

func TestGetHostId_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		HostIncreaseId: 0,
	}

	id1 := db.GetHostId()
	id2 := db.GetHostId()

	if id1 != 1 {
		t.Errorf("Expected 1, got %d", id1)
	}

	if id2 != 2 {
		t.Errorf("Expected 2, got %d", id2)
	}
}
