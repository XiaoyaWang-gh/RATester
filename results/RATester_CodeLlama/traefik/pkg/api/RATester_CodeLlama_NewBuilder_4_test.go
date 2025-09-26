package api

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewBuilder_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	staticConfig := static.Configuration{}
	builder := NewBuilder(staticConfig)
	assert.NotNil(t, builder)
}
