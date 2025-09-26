package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testStruct := &StreamExampleReply{
		Data: "testData",
	}

	testStruct.Reset()

	if testStruct.Data != "" {
		t.Errorf("Expected Data to be empty after Reset, got %s", testStruct.Data)
	}
}
