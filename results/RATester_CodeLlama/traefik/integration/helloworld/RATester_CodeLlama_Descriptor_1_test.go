package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestDescriptor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reply := &StreamExampleReply{}
	reply.Descriptor()
}
