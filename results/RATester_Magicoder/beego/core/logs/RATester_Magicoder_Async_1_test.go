package logs

import (
	"fmt"
	"testing"
)

func TestAsync_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	msgLen := []int64{100}
	result := bl.Async(msgLen...)
	if result != bl {
		t.Errorf("Expected %v, got %v", bl, result)
	}
}
