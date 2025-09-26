package try

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestRequestWithTransport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := &http.Request{}
	timeout := time.Duration(10)
	transport := &http.Transport{}
	conditions := []ResponseCondition{}

	// when
	err := RequestWithTransport(req, timeout, transport, conditions...)

	// then
	assert.NoError(t, err)
}
