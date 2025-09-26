package acme

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &ChallengeHTTP{}
	// when
	timeout, interval := c.Timeout()
	// then
	assert.Equal(t, 60*time.Second, timeout)
	assert.Equal(t, 5*time.Second, interval)
}
