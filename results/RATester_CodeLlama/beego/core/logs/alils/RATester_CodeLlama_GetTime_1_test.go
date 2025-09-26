package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Log{}
	m.Time = proto.Uint32(100)
	if m.GetTime() != 100 {
		t.Errorf("GetTime() = %v, want %v", m.GetTime(), 100)
	}
}
