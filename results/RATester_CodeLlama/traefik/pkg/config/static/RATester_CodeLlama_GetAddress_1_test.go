package static

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetAddress_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ep := EntryPoint{
		Address: "127.0.0.1:8080",
	}
	assert.Equal(t, "127.0.0.1:8080", ep.GetAddress())
}
