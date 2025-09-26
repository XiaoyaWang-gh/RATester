package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogContent{}
	m.Key = proto.String("key")
	m.Value = proto.String("value")
	if m.GetValue() != "value" {
		t.Errorf("GetValue() = %v, want %v", m.GetValue(), "value")
	}
}
