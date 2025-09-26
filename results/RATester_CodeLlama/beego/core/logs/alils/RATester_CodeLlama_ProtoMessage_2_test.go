package alils

import (
	"fmt"
	"testing"
)

func TestProtoMessage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m LogGroupList
	m.ProtoMessage()
}
