package consulcatalog

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRepeatSend_1(t *testing.T) {
	t.Run("test repeatSend", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		c := make(chan struct{}, 1)
		go repeatSend(ctx, time.Second, c)

		select {
		case <-ctx.Done():
			t.Error("context should not be done")
		case <-c:
			t.Log("received value from channel")
		}

		cancel()
		select {
		case <-ctx.Done():
			t.Log("context should be done")
		case <-c:
			t.Error("channel should be closed")
		}
	})
}
