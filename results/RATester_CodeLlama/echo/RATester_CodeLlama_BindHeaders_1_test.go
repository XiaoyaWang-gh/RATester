package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBindHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	b := &DefaultBinder{}
	c := &context{}
	i := &struct {
		Header string `header:"header"`
	}{}

	// when
	err := b.BindHeaders(c, i)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "header", i.Header)
}
