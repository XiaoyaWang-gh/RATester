package fiber

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var z redirectionMsgs
	var en *msgp.Writer
	var err error

	en = msgp.NewWriter(nil)
	err = z.EncodeMsg(en)
	if err != nil {
		t.Fatal(err)
	}
	en.Flush()
}
