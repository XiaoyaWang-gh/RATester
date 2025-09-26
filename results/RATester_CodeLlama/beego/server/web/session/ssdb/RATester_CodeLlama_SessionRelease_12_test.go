package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRelease_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionStore{
		sid: "123456",
	}
	s.SessionRelease(context.Background(), nil)
}
