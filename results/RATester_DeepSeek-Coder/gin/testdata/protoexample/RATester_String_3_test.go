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
		Type:  proto.Int32(77),
		Reps:  []int64{1, 2, 3},
	}

	expected := `label:"testLabel" type:77 reps:1 reps:2 reps:3`
	result := test.String()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
