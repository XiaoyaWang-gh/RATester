package warpc

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestGetDispatcher_1(t *testing.T) {
	t.Parallel()

	type Q = int
	type R = string

	pool := &dispatcherPool[Q, R]{
		counter:     atomic.Uint32{},
		dispatchers: []*dispatcher[Q, R]{},
		close:       nil,
		errc:        make(chan error),
		donec:       make(chan struct{}),
	}

	dispatcher1 := &dispatcher[Q, R]{}
	dispatcher2 := &dispatcher[Q, R]{}
	pool.dispatchers = append(pool.dispatchers, dispatcher1, dispatcher2)

	t.Run("getDispatcher", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := pool.getDispatcher()
		want := dispatcher1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("getDispatcher", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := pool.getDispatcher()
		want := dispatcher2
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("getDispatcher", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := pool.getDispatcher()
		want := dispatcher1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
