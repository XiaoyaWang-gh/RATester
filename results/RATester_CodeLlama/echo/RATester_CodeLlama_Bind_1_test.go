package echo

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/zeebo/assert"
)

func TestBind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	b := &DefaultBinder{}
	i := &struct {
		Name string `query:"name"`
	}{}
	c := &context{
		query: url.Values{
			"name": []string{"john"},
		},
	}

	// when
	err := b.Bind(i, c)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "john", i.Name)
}
