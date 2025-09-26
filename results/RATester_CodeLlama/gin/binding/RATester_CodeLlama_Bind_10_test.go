package binding

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := &http.Request{
		Body: ioutil.NopCloser(bytes.NewBufferString("")),
	}
	obj := &struct{}{}

	// when
	err := msgpackBinding{}.Bind(req, obj)

	// then
	assert.NoError(t, err)
}
