package web

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXSRFFormHTML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.EnableXSRF = true
	c.XSRFExpire = 10
	c.XSRFToken()
	assert.Equal(t, `<input type="hidden" name="_xsrf" value="`+c.XSRFToken()+`" />`, c.XSRFFormHTML())
}
