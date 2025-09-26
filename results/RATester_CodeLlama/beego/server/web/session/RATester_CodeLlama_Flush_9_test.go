package session

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &FileSessionStore{}
	fs.Flush(context.Background())
}
