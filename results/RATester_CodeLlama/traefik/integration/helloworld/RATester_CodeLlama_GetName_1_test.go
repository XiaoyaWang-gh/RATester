package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &HelloRequest{Name: "test"}
	if got := m.GetName(); got != "test" {
		t.Errorf("HelloRequest.GetName() = %v, want %v", got, "test")
	}
}
