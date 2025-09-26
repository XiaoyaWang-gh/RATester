package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Label: proto.String("testLabel"),
		Type:  proto.Int32(1),
		Reps:  []int64{1, 2, 3},
	}

	expected := "label: \"testLabel\"\n" +
		"type: 1\n" +
		"reps: 1\n" +
		"reps: 2\n" +
		"reps: 3\n"

	if test.String() != expected {
		t.Errorf("Test.String() = %s; want %s", test.String(), expected)
	}
}
