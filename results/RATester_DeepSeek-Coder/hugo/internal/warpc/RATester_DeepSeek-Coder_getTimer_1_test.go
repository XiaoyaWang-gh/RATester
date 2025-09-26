package warpc

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimer_1(t *testing.T) {
	t.Run("should return a new Timer if the pool is empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		want := time.NewTimer(1 * time.Second)
		got := getTimer(1 * time.Second)
		if got.C != want.C {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})

	t.Run("should return a Timer from the pool if it's not empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		want := time.NewTimer(1 * time.Second)
		timerPool.Put(want)
		got := getTimer(1 * time.Second)
		if got.C != want.C {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})

	t.Run("should reset the Timer if it's not empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		want := time.NewTimer(1 * time.Second)
		timerPool.Put(want)
		getTimer(2 * time.Second)
		if !want.Stop() {
			<-want.C
		}
		got := getTimer(1 * time.Second)
		if got.C == want.C {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}
