package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/go-acme/lego/v4/lego"
)

func TestSetDefaults_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Configuration{}
	a.SetDefaults()
	assert.Equal(t, lego.LEDirectoryProduction, a.CAServer)
	assert.Equal(t, "acme.json", a.Storage)
	assert.Equal(t, "RSA4096", a.KeyType)
	assert.Equal(t, 3*30*24, a.CertificatesDuration)
}
