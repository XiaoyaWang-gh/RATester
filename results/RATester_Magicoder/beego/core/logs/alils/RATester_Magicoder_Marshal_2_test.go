package alils

import (
	"fmt"
	"testing"
)

func TestMarshal_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "testKey"
	value := "testValue"
	logContent := &LogContent{
		Key:   &key,
		Value: &value,
	}
	data, err := logContent.Marshal()
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	if len(data) == 0 {
		t.Errorf("Marshal() returned empty data")
	}
}
