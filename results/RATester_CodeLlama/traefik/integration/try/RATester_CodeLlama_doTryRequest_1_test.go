package try

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestDoTryRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	request := &http.Request{}
	timeout := time.Duration(10)
	transport := &http.Transport{}
	conditions := []ResponseCondition{}

	// when
	_, err := doTryRequest(request, timeout, transport, conditions...)

	// then
	assert.NoError(t, err)
}
