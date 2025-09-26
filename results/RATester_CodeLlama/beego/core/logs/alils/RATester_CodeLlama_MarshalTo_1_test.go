package alils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestMarshalTo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m LogGroup
	m.Logs = append(m.Logs, &Log{
		Time: proto.Uint32(1234567890),
		Contents: []*LogContent{
			{
				Key:   proto.String("key1"),
				Value: proto.String("value1"),
			},
			{
				Key:   proto.String("key2"),
				Value: proto.String("value2"),
			},
		},
	})
	m.Reserved = proto.String("reserved")
	m.Topic = proto.String("topic")
	m.Source = proto.String("source")

	data, err := m.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	var m2 LogGroup
	if err := m2.Unmarshal(data); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(&m, &m2) {
		t.Fatalf("m1 != m2")
	}
}
