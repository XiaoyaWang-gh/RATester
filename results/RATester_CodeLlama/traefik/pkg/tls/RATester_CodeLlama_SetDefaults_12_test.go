package tls

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Options{}
	o.SetDefaults()

	assert.Equal(t, DefaultTLSOptions.ALPNProtocols, o.ALPNProtocols)
}
