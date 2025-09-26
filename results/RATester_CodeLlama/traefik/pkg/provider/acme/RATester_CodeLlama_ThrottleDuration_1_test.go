package acme

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestThrottleDuration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	assert.Equal(t, time.Duration(0), p.ThrottleDuration())
}
