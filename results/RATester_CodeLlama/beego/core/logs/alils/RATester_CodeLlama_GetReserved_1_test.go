package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetReserved_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogGroup{}
	m.Reserved = proto.String("test")
	if m.GetReserved() != "test" {
		t.Errorf("GetReserved() = %v, want %v", m.GetReserved(), "test")
	}
}
