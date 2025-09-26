package mock

import (
	"fmt"
	"testing"
)

func TestStartMock_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mock := StartMock()
	if mock == nil {
		t.Errorf("StartMock() = %v, want not nil", mock)
	}
}
