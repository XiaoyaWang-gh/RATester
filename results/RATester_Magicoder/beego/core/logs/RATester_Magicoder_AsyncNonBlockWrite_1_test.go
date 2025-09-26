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

	bl := &BeeLogger{
		asynchronous: true,
	}

	result := bl.AsyncNonBlockWrite()

	if result.logWithNonBlocking != true {
		t.Error("Expected logWithNonBlocking to be true, but got false")
	}
}
