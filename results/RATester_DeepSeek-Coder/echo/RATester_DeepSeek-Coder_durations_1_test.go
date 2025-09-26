package echo

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestDurations_1(t *testing.T) {
	t.Run("Test durations with valid duration strings", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{}
		values := []string{"1s", "2m", "3h"}
		dest := []time.Duration{}
		b.durations("sourceParam", values, &dest)
		assert.Nil(t, b.errors)
		assert.Equal(t, 3, len(dest))
		assert.Equal(t, time.Second, dest[0])
		assert.Equal(t, 2*time.Minute, dest[1])
		assert.Equal(t, 3*time.Hour, dest[2])
	})

	t.Run("Test durations with invalid duration strings", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{}
		values := []string{"1s", "invalid", "3h"}
		dest := []time.Duration{}
		b.durations("sourceParam", values, &dest)
		assert.NotNil(t, b.errors)
		assert.Equal(t, 1, len(b.errors))
		assert.Equal(t, 2, len(dest))
		assert.Equal(t, time.Second, dest[0])
		assert.Equal(t, 3*time.Hour, dest[1])
	})

	t.Run("Test durations with failFast set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{failFast: true}
		values := []string{"1s", "invalid", "3h"}
		dest := []time.Duration{}
		b.durations("sourceParam", values, &dest)
		assert.NotNil(t, b.errors)
		assert.Equal(t, 1, len(b.errors))
		assert.Equal(t, 1, len(dest))
		assert.Equal(t, time.Second, dest[0])
	})
}
