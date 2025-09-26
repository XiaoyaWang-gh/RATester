package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestBuildMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{
		ResolverName: "test",
	}
	conf := p.buildMessage()
	assert.Equal(t, "test.acme", conf.ProviderName)
	assert.NotNil(t, conf.Configuration)
	assert.NotNil(t, conf.Configuration.HTTP)
	assert.NotNil(t, conf.Configuration.TLS)
}
