package ip

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains_2(t *testing.T) {
	ip := &Checker{}

	t.Run("empty IP address", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		contains, err := ip.Contains("")
		require.Error(t, err)
		require.False(t, contains)
	})

	t.Run("unable to parse address", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		contains, err := ip.Contains("127.0.0.1")
		require.Error(t, err)
		require.False(t, contains)
	})

	t.Run("address not in trusted IPs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ip.authorizedIPs = []*net.IP{
			{},
		}

		contains, err := ip.Contains("127.0.0.1")
		require.NoError(t, err)
		require.False(t, contains)
	})

	t.Run("address in trusted IPs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ip.authorizedIPs = []*net.IP{
			{},
		}

		contains, err := ip.Contains("127.0.0.1")
		require.NoError(t, err)
		require.True(t, contains)
	})
}
