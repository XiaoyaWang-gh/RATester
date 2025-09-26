package hugocontext

import (
	"fmt"
	"testing"
)

func TestDump_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &HugoContext{
		Pid: 123,
	}
	ctx.Dump([]byte("test"), 0)
}
