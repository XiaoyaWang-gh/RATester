package csrf

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalMsg_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := storageManager{}
	b := make([]byte, 0)
	expectedOutput := []byte{0x80}
	output, err := manager.MarshalMsg(b)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}
