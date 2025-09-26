package proxy

import (
	"fmt"
	"testing"
)

func TestGetUsedPortsNum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		usedPortsNum: 10,
	}

	result := pxy.GetUsedPortsNum()

	if result != 10 {
		t.Errorf("Expected 10, but got %d", result)
	}
}
