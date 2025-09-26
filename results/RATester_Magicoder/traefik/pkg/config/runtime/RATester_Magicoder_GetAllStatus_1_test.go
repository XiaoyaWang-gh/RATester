package runtime

import (
	"fmt"
	"testing"
)

func TestGetAllStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serviceInfo := &ServiceInfo{
		serverStatus: map[string]string{
			"server1": "status1",
			"server2": "status2",
		},
	}

	allStatus := serviceInfo.GetAllStatus()

	if len(allStatus) != 2 {
		t.Errorf("Expected 2 servers, got %d", len(allStatus))
	}

	if allStatus["server1"] != "status1" {
		t.Errorf("Expected server1 status to be 'status1', got %s", allStatus["server1"])
	}

	if allStatus["server2"] != "status2" {
		t.Errorf("Expected server2 status to be 'status2', got %s", allStatus["server2"])
	}
}
