package fmt

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_25(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ns := New(d)

	if ns.logger != d.Log {
		t.Errorf("Expected logger to be %v but was %v", d.Log, ns.logger)
	}
}
