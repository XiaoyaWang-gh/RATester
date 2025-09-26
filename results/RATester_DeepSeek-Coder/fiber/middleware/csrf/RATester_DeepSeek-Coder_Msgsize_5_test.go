package csrf

import (
	"fmt"
	"testing"
)

func TestMsgsize_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &storageManager{}
	size := manager.Msgsize()
	if size != 1 {
		t.Errorf("Expected size to be 1, got %d", size)
	}
}
