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
		t.Error("Expected mock to not be nil")
	}

	mock.Mock(nil, nil, nil)
	mock.Clear()
	mock.MockByPath("", nil, nil)
}
