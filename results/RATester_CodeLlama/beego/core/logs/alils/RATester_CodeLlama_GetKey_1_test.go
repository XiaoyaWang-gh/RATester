package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogContent{}
	m.Key = proto.String("key")
	m.Value = proto.String("value")
	if m.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", m.GetKey(), "key")
	}
}
