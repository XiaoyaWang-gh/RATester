package echo

import (
	"fmt"
	"testing"
)

func TestHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	hostName := "localhost"

	g := e.Host(hostName)

	if g.host != hostName {
		t.Errorf("Expected host name to be %s, got %s", hostName, g.host)
	}

	if g.echo != e {
		t.Errorf("Expected echo instance to be %v, got %v", e, g.echo)
	}
}
