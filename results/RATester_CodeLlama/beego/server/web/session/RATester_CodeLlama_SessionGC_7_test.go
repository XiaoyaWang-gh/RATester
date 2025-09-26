package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pder := &CookieProvider{}
	ctx := context.Background()
	pder.SessionGC(ctx)
}
