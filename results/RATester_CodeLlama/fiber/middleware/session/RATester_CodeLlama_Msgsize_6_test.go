package session

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMsgsize_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var z *data
	var s int
	s = 1 + 5 + msgp.MapHeaderSize
	if z.Data != nil {
		for za0001, za0002 := range z.Data {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.GuessSize(za0002)
		}
	}
	if s != z.Msgsize() {
		t.Logf("size of z.Data is not equal to z.Msgsize()")
		t.Fail()
	}
}
