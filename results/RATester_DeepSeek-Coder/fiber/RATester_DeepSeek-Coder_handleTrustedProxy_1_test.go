package fiber

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestHandleTrustedProxy_1(t *testing.T) {
	t.Parallel()

	app := &App{
		config: Config{
			trustedProxiesMap: make(map[string]struct{}),
		},
	}

	t.Run("single IP", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ip := "192.0.2.0"
		app.handleTrustedProxy(ip)

		if _, ok := app.config.trustedProxiesMap[ip]; !ok {
			t.Errorf("Expected IP %s to be in the trusted proxies map", ip)
		}
	})

	t.Run("CIDR IP", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ip := "192.0.2.0/24"
		app.handleTrustedProxy(ip)

		if len(app.config.trustedProxyRanges) != 1 {
			t.Errorf("Expected 1 trusted proxy range, got %d", len(app.config.trustedProxyRanges))
		}

		_, ipNet, err := net.ParseCIDR(ip)
		if err != nil {
			t.Errorf("Failed to parse CIDR IP: %v", err)
		}

		if !reflect.DeepEqual(app.config.trustedProxyRanges[0], ipNet) {
			t.Errorf("Expected trusted proxy range to be %s, got %s", ipNet, app.config.trustedProxyRanges[0])
		}
	})

	t.Run("invalid CIDR IP", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ip := "192.0.2.0/33"
		app.handleTrustedProxy(ip)

		if len(app.config.trustedProxyRanges) != 1 {
			t.Errorf("Expected 1 trusted proxy range, got %d", len(app.config.trustedProxyRanges))
		}
	})
}
