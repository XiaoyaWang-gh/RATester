package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewIPChecker_1(t *testing.T) {
	t.Run("Test with default options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		checker := newIPChecker(nil)
		assert.True(t, checker.trustLoopback)
		assert.True(t, checker.trustLinkLocal)
		assert.True(t, checker.trustPrivateNet)
	})

	t.Run("Test with custom options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		checker := newIPChecker([]TrustOption{
			func(c *ipChecker) {
				c.trustLoopback = false
			},
			func(c *ipChecker) {
				c.trustLinkLocal = false
			},
			func(c *ipChecker) {
				c.trustPrivateNet = false
			},
		})
		assert.False(t, checker.trustLoopback)
		assert.False(t, checker.trustLinkLocal)
		assert.False(t, checker.trustPrivateNet)
	})
}
