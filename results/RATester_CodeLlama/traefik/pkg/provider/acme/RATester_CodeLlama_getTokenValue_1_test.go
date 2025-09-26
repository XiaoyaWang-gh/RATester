package acme

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetTokenValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &ChallengeHTTP{
		httpChallenges: map[string]map[string][]byte{
			"token": {
				"domain": []byte("value"),
			},
		},
	}
	ctx := context.Background()
	token := "token"
	domain := "domain"

	// when
	result := c.getTokenValue(ctx, token, domain)

	// then
	assert.Equal(t, []byte("value"), result)
}
