package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pder := &MemProvider{
		maxlifetime: 10,
	}
	pder.SessionInit(context.Background(), 10, "")
	pder.SessionRegenerate(context.Background(), "", "123456")
	pder.SessionGC(context.Background())
}
