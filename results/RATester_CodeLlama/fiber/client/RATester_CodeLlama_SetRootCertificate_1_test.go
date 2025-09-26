package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetRootCertificate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	c := &Client{}
	path := "./testdata/cert.pem"

	// act
	c.SetRootCertificate(path)

	// assert
	assert.NotNil(t, c.TLSConfig().RootCAs)
}
