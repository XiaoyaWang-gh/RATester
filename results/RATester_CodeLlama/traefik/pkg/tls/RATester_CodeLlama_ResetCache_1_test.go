package tls

import (
	"fmt"
	"testing"
)

func TestResetCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := CertificateStore{}
	c.ResetCache()
}
