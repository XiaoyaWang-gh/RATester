package deps

import (
	"fmt"
	"testing"
)

func TestNotify_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Listeners{}
	b.Add(func() {
		t.Log("called")
	})
	b.Notify()
}
