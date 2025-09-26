package try

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	url := "http://www.example.com"
	timeout := time.Duration(10)
	conditions := []ResponseCondition{
		func(resp *http.Response) error {
			return nil
		},
	}

	// when
	err := GetRequest(url, timeout, conditions...)

	// then
	require.NoError(t, err)
}
