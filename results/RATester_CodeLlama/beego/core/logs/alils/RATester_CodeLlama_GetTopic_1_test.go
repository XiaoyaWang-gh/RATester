package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetTopic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogGroup{}
	m.Topic = proto.String("test")
	if m.GetTopic() != "test" {
		t.Errorf("GetTopic() = %v, want %v", m.GetTopic(), "test")
	}
}
