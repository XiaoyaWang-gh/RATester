package mock

import (
	"errors"
	"fmt"
	"testing"
)

func TestMockRollback_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mock := MockRollback(errors.New("mock error"))
	if mock == nil {
		t.Error("MockRollback returned nil")
	}
	if mock.cond == nil {
		t.Error("MockRollback returned mock with nil condition")
	}
	if len(mock.resp) != 1 {
		t.Errorf("MockRollback returned mock with unexpected response length: %d", len(mock.resp))
	}
	if mock.resp[0] != errors.New("mock error") {
		t.Errorf("MockRollback returned mock with unexpected error: %v", mock.resp[0])
	}
	if mock.cb != nil {
		t.Error("MockRollback returned mock with non-nil callback")
	}
}
