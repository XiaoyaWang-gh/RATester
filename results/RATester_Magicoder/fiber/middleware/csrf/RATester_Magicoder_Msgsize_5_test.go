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

	manager := storageManager{}
	expected := 1
	actual := manager.Msgsize()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
