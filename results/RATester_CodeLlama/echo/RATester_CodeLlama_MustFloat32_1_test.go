package echo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustFloat32_1(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := new(float32)

		b := new(ValueBinder)
		b.ValueFunc = func(sourceParam string) string {
			return "1.23"
		}

		actual := b.MustFloat32(sourceParam, dest)

		assert.Equal(t, b, actual)
		assert.Equal(t, float32(1.23), *dest)
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := new(float32)

		b := new(ValueBinder)
		b.ValueFunc = func(sourceParam string) string {
			return ""
		}

		assert.PanicsWithError(t, `echo: parameter "sourceParam" does not exist`, func() {
			b.MustFloat32(sourceParam, dest)
		})
	})
}
