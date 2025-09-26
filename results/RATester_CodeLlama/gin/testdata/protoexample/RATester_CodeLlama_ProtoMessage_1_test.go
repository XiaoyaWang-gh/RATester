package protoexample

import (
	"fmt"
	"testing"
)

func TestProtoMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var test Test_OptionalGroup
	test.ProtoMessage()
}
