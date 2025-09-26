package virtual

import (
	"fmt"
	"testing"
)

func TestService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	if c.Service() != c.svr {
		t.Errorf("Service() = %v, want %v", c.Service(), c.svr)
	}
}
