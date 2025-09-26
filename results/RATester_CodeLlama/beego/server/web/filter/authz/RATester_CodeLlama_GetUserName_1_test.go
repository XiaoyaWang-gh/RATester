package authz

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestGetUserName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	a := &BasicAuthorizer{}
	r := &http.Request{}
	// when
	username := a.GetUserName(r)
	// then
	assert.Equal(t, "", username)
}
