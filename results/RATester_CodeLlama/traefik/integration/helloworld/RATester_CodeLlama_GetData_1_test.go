package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &StreamExampleReply{}
	if got := m.GetData(); got != "" {
		t.Errorf("StreamExampleReply.GetData() = %v, want %v", got, "")
	}
}
