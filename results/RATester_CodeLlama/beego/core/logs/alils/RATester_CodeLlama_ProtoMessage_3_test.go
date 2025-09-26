package alils

import (
	"fmt"
	"testing"
)

func TestProtoMessage_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	log := &Log{}
	log.ProtoMessage()
}
