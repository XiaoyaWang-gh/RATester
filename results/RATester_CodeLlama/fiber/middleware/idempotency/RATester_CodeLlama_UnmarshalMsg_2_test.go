package idempotency

import (
	"bytes"
	"fmt"
	"testing"
)

func TestUnmarshalMsg_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := response{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalMsg(bts[0:0])
	bts, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	bts, err = v.MarshalMsg(bts[0:0])
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts, []byte{}) {
		t.Fatal("Wrong msgsize")
	}
}
