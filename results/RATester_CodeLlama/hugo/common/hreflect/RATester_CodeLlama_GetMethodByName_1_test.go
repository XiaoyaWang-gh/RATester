package hreflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetMethodByName_1(t *testing.T) {
	t.Parallel()

	t.Run("should return a method value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		v := reflect.ValueOf(struct{}{})
		name := "Method"

		// when
		actual := GetMethodByName(v, name)

		// then
		assert.NotEmpty(t, actual)
	})

	t.Run("should return an empty value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		v := reflect.ValueOf(struct{}{})
		name := "NotExist"

		// when
		actual := GetMethodByName(v, name)

		// then
		assert.Empty(t, actual)
	})
}
