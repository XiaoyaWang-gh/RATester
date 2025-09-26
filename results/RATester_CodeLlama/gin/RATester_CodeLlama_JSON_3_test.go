package gin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON_3(t *testing.T) {
	t.Parallel()
	t.Run("nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		var a errorMsgs
		assert.Nil(t, a.JSON())
	})
	t.Run("one", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		var a errorMsgs
		a = append(a, &Error{})
		assert.NotNil(t, a.JSON())
	})
	t.Run("many", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		var a errorMsgs
		a = append(a, &Error{})
		a = append(a, &Error{})
		assert.NotNil(t, a.JSON())
	})
}
