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

	log := &Log{
		Time: proto.Uint32(123),
	}

	time := log.GetTime()

	if time != 123 {
		t.Errorf("Expected time to be 123, but got %d", time)
	}
}
