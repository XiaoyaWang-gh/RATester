package binding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindBody_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	body := []byte(`{"name":"test"}`)
	obj := &struct {
		Name string `json:"name"`
	}{}
	// when
	err := jsonBinding{}.BindBody(body, obj)
	// then
	assert.NoError(t, err)
	assert.Equal(t, "test", obj.Name)
}
