package testhelpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAdd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &CollectingGauge{}
	g.Add(1)
	assert.Equal(t, float64(1), g.GaugeValue)
}
