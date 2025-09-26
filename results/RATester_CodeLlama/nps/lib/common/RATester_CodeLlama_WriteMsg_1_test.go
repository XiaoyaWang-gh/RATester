package common

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteMsg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &StoreMsg{}
	lg.Init("")
	lg.WriteMsg(time.Now(), "test", 1)
	lg.Flush()
	lg.Destroy()
}
