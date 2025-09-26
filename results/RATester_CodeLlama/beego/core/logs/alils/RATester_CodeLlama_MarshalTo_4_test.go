package alils

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestMarshalTo_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m LogContent
	m.Key = proto.String("key")
	m.Value = proto.String("value")
	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal() = %v", err)
	}
	size := m.Size()
	if len(data) != size {
		t.Errorf("Size() = %v, but marshaled size = %v", size, len(data))
	}
	data2 := make([]byte, size)
	n, err := m.MarshalTo(data2)
	if err != nil {
		t.Fatalf("MarshalTo() = %v", err)
	}
	if n != size {
		t.Errorf("MarshalTo() = %v, but marshaled size = %v", n, size)
	}
	if !bytes.Equal(data, data2) {
		t.Errorf("MarshalTo() = %v, but marshaled data = %v", data2, data)
	}
}
