package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestGetSource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogGroup{}
	m.Source = proto.String("test")
	if m.GetSource() != "test" {
		t.Errorf("GetSource() = %v, want %v", m.GetSource(), "test")
	}
}
