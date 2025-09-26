package consulcatalog

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestLeafWatcherHandler_1(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dest := make(chan keyPair, 1)
	handler := leafWatcherHandler(ctx, dest)

	t.Run("nil input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler(nil, nil)
		select {
		case <-dest:
			t.Error("Expected no update")
		default:
		}
	})

	t.Run("invalid input type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler(nil, "invalid")
		select {
		case <-dest:
			t.Error("Expected no update")
		default:
		}
	})

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cert := "test-cert"
		key := "test-key"
		handler(nil, &api.LeafCert{CertPEM: cert, PrivateKeyPEM: key})
		select {
		case kp := <-dest:
			if kp.cert != cert || kp.key != key {
				t.Errorf("Expected keyPair{%s, %s}, got keyPair{%s, %s}", cert, key, kp.cert, kp.key)
			}
		default:
			t.Error("Expected update")
		}
	})
}
