package tls

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewCertificateStore_1(t *testing.T) {
	t.Run("should return a new CertificateStore", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store := NewCertificateStore()

		assert.NotNil(t, store)
		assert.NotNil(t, store.DynamicCerts)
		assert.NotNil(t, store.CertCache)
	})

	t.Run("should return a new Safe", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store := NewCertificateStore()

		assert.NotNil(t, store.DynamicCerts)
	})

	t.Run("should return a new Cache", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store := NewCertificateStore()

		assert.NotNil(t, store.CertCache)
	})
}
