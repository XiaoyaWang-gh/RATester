package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	x := &Test_OptionalGroup{
		RequiredField: proto.String("hello"),
	}
	want := "RequiredField: \"hello\""
	if got := x.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
