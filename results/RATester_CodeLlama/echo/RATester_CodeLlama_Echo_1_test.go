package echo

import (
	"fmt"
	"testing"
)

func TestEcho_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		echo: &Echo{},
	}
	if c.Echo() != c.echo {
		t.Errorf("Echo() = %v, want %v", c.Echo(), c.echo)
	}
}
