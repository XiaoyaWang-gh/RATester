package debug

import (
	"fmt"
	"testing"
)

func TestTimer_1(t *testing.T) {
	ns := &Namespace{}
	t.Run("nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		if got := ns.Timer(name); got != nopTimer {
			t.Errorf("Timer() = %v, want %v", got, nopTimer)
		}
	})
}
