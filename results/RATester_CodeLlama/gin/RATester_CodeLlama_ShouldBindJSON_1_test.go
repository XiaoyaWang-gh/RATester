package gin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBindJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// given
	c := &Context{}
	c.Request = &http.Request{}
	c.Request.Body = ioutil.NopCloser(bytes.NewBufferString(`{"name":"test"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// when
	err := c.ShouldBindJSON(&struct{ Name string }{""})

	// then
	assert.NoError(t, err)
}
