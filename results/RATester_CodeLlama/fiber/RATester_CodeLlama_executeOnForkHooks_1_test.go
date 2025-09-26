package fiber

import (
	"fmt"
	"testing"
)

func TestExecuteOnForkHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	h := &Hooks{}
	h.OnFork(func(pid int) error {
		if pid != 1 {
			t.Errorf("pid should be 1, but got %d", pid)
		}
		return nil
	})
	// when
	h.executeOnForkHooks(1)
	// then
}
