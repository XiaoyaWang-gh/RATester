package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestCheckFlowAndConnNum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &BaseServer{}
	client := &file.Client{
		Flow: &file.Flow{
			FlowLimit:  100,
			ExportFlow: 100,
			InletFlow:  100,
		},
	}

	// Test case 1: Traffic exceeded
	client.Flow.ExportFlow = 200
	client.Flow.InletFlow = 100
	err := server.CheckFlowAndConnNum(client)
	if err == nil {
		t.Error("Expected error for traffic exceeded, but got nil")
	}

	// Test case 2: Connections exceed the current client limit
	client.Flow.ExportFlow = 100
	client.Flow.InletFlow = 100
	client.NowConn = 10
	err = server.CheckFlowAndConnNum(client)
	if err == nil {
		t.Error("Expected error for connections exceed the current client limit, but got nil")
	}

	// Test case 3: No error
	client.Flow.ExportFlow = 100
	client.Flow.InletFlow = 100
	client.NowConn = 9
	err = server.CheckFlowAndConnNum(client)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
