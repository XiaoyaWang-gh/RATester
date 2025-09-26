package protoexample

import (
	"fmt"
	"testing"
)

func Testinit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file_test_proto_init()
}
