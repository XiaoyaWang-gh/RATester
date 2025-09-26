package acme

import (
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/lego"
)

func TestSetDefaults_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Configuration{}
	config.SetDefaults()

	if config.CAServer != lego.LEDirectoryProduction {
		t.Errorf("Expected CAServer to be %s, but got %s", lego.LEDirectoryProduction, config.CAServer)
	}

	if config.Storage != "acme.json" {
		t.Errorf("Expected Storage to be %s, but got %s", "acme.json", config.Storage)
	}

	if config.KeyType != "RSA4096" {
		t.Errorf("Expected KeyType to be %s, but got %s", "RSA4096", config.KeyType)
	}
}
