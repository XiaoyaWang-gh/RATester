package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustUint16s_1(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := []uint16{}

		b := &ValueBinder{}
		b.MustUint16s(sourceParam, &dest)

		assert.Equal(t, []uint16{}, dest)
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := []uint16{}

		b := &ValueBinder{}
		b.MustUint16s(sourceParam, &dest)

		assert.Equal(t, []uint16{}, dest)
	})
}
