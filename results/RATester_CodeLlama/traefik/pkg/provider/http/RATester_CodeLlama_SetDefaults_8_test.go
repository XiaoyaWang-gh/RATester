package http

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	ptypes "github.com/traefik/paerser/types"
)

func TestSetDefaults_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.SetDefaults()
	assert.Equal(t, ptypes.Duration(5*time.Second), p.PollInterval)
	assert.Equal(t, ptypes.Duration(5*time.Second), p.PollTimeout)
}
