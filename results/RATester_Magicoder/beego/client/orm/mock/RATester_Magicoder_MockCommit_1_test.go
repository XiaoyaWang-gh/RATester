package mock

import (
	"fmt"
	"testing"
)

func TestMockCommit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mock := MockCommit(nil)
	if mock == nil {
		t.Error("MockCommit returned nil")
	}
}
