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

	bl.AsyncNonBlockWrite()

	if !bl.logWithNonBlocking {
		t.Errorf("Expected logWithNonBlocking to be true, but got false")
	}
}
