package logs

import (
	"fmt"
	"testing"
)

func TestSetLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	adapterName := "file"
	configs := []string{"test.log"}
	err := bl.SetLogger(adapterName, configs...)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
