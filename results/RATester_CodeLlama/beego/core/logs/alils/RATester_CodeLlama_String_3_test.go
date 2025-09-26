package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogContent{}
	m.Key = proto.String("key")
	m.Value = proto.String("value")
	if m.String() != "key: \"value\"" {
		t.Errorf("String() = %v, want %v", m.String(), "key: \"value\"")
	}
}
