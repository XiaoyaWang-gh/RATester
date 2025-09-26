package proxy

import (
	"fmt"
	"testing"
)

func TestBalancerForward_1(t *testing.T) {
	t.Parallel()
	t.Run("TestBalancerForward", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		// TODO
	})
}
