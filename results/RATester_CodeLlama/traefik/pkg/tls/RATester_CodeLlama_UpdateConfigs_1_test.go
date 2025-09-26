package tls

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestUpdateConfigs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	m := &Manager{}

	stores := map[string]Store{
		"default": {
			DefaultCertificate: &Certificate{
				CertFile: "certFile",
				KeyFile:  "keyFile",
			},
		},
	}

	configs := map[string]Options{
		"default": {
			MinVersion: "VersionTLS12",
		},
	}

	certs := []*CertAndStores{
		{
			Certificate: Certificate{
				CertFile: "certFile",
				KeyFile:  "keyFile",
			},
			Stores: []string{"default"},
		},
	}

	m.UpdateConfigs(ctx, stores, configs, certs)

	assert.Equal(t, 1, len(m.stores))
	assert.Equal(t, 1, len(m.storesConfig))
	assert.Equal(t, 1, len(m.configs))
	assert.Equal(t, 1, len(m.certs))
}
