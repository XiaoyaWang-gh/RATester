package ping

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Handler{}
	h.SetDefaults()
	assert.Equal(t, "traefik", h.EntryPoint)
	assert.Equal(t, http.StatusServiceUnavailable, h.TerminatingStatusCode)
}
