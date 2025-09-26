package try

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := &http.Request{}
	timeout := time.Duration(10)
	conditions := []ResponseCondition{}

	// when
	err := Request(req, timeout, conditions...)

	// then
	require.NoError(t, err)
}
