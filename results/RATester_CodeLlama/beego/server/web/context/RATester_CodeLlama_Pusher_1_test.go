package context

import (
	"fmt"
	"testing"
)

func TestPusher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	pusher := r.Pusher()
	if pusher != nil {
		t.Errorf("Pusher() = %v, want nil", pusher)
	}
}
