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

	ctx := &context{
		echo: &Echo{},
	}

	expected := ctx.echo
	actual := ctx.Echo()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
