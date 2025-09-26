package common

import (
	"fmt"
	"testing"
)

func TestFlush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storeMsg := &StoreMsg{}
	storeMsg.Flush()
}
