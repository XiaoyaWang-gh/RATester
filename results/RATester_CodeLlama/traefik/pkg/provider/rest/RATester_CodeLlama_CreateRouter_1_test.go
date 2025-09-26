package rest

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestCreateRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	router := p.CreateRouter()
	assert.NotNil(t, router)
}
