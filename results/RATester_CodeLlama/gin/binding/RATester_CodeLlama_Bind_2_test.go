package binding

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind_2(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		obj := &struct{}{}

		// when
		err := protobufBinding{}.Bind(req, obj)

		// then
		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		obj := &struct{}{}

		// when
		err := protobufBinding{}.Bind(req, obj)

		// then
		assert.Error(t, err)
	})
}
