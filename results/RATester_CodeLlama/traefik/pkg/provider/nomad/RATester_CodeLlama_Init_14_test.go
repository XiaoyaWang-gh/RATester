package nomad

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestInit_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.Init()
	assert.Equal(t, providerName, p.name)
}
