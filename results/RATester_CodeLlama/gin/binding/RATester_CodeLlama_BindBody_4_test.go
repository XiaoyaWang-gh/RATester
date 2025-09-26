package binding

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/require"
)

func TestBindBody_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	body := []byte("<xml><name>John</name><age>25</age></xml>")
	obj := &struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}{}

	// when
	err := xmlBinding{}.BindBody(body, obj)

	// then
	require.NoError(t, err)
	assert.Equal(t, "John", obj.Name)
	assert.Equal(t, 25, obj.Age)
}
