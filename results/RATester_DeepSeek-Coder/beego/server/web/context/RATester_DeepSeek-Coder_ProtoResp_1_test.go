package context

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProtoResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		proto.Message
		Field1 string `protobuf:"bytes,1,opt,name=field1,proto3" json:"field1,omitempty"`
		Field2 int32  `protobuf:"varint,2,opt,name=field2,proto3" json:"field2,omitempty"`
	}

	testStruct := &TestStruct{
		Field1: "test",
		Field2: 123,
	}

	ctx := &Context{
		Output: &BeegoOutput{},
	}

	err := ctx.ProtoResp(testStruct)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
