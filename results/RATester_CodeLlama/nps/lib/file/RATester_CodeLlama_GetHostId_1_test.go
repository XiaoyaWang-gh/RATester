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

	//given
	var jsonDb JsonDb
	jsonDb.HostIncreaseId = 100
	//when
	hostId := jsonDb.GetHostId()
	//then
	if hostId != 101 {
		t.Errorf("GetHostId() = %v, want %v", hostId, 101)
	}
}
