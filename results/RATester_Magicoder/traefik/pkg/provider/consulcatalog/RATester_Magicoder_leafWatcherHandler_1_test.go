package consulcatalog

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestLeafWatcherHandler_1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dest := make(chan keyPair, 1)
	handler := leafWatcherHandler(ctx, dest)

	t.Run("nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler(nil, nil)
		select {
		case <-dest:
			t.Error("expected no value to be sent on dest")
		default:
		}
	})

	t.Run("invalid type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler(nil, "invalid")
		select {
		case <-dest:
			t.Error("expected no value to be sent on dest")
		default:
		}
	})

	t.Run("valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cert := "cert"
		key := "key"
		handler(nil, &api.LeafCert{CertPEM: cert, PrivateKeyPEM: key})
		select {
		case kp := <-dest:
			if kp.cert != cert || kp.key != key {
				t.Errorf("expected %v, got %v", keyPair{cert, key}, kp)
			}
		default:
			t.Error("expected value to be sent on dest")
		}
	})
}
