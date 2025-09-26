package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &StreamExampleReply{}
	m.Reset()
	if m.GetData() != "" {
		t.Errorf("m.GetData() = %v, want %v", m.GetData(), "")
	}
}
