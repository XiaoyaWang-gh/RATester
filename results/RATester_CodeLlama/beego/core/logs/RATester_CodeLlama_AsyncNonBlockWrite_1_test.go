package logs

import (
	"fmt"
	"testing"
)

func TestAsyncNonBlockWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	bl.AsyncNonBlockWrite()
	if bl.logWithNonBlocking != true {
		t.Error("AsyncNonBlockWrite failed")
	}
}
