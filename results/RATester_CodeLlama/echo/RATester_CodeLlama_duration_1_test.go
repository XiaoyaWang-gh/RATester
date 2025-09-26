package echo

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestDuration_1(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := new(time.Duration)
		valueMustExist := true

		b := new(ValueBinder)
		b.ValueFunc = func(sourceParam string) string {
			return "10s"
		}

		b.duration(sourceParam, dest, valueMustExist)

		assert.Equal(t, time.Duration(10*time.Second), *dest)
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := new(time.Duration)
		valueMustExist := true

		b := new(ValueBinder)
		b.ValueFunc = func(sourceParam string) string {
			return ""
		}

		b.duration(sourceParam, dest, valueMustExist)

		assert.Equal(t, (*time.Duration)(nil), dest)
	})
}
