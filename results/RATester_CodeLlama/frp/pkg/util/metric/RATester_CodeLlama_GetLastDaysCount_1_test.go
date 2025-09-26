package metric

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetLastDaysCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &StandardDateCounter{
		reserveDays: 10,
		counts:      make([]int64, 10),
	}
	c.counts[0] = 1
	c.counts[1] = 2
	c.counts[2] = 3
	c.counts[3] = 4
	c.counts[4] = 5
	c.counts[5] = 6
	c.counts[6] = 7
	c.counts[7] = 8
	c.counts[8] = 9
	c.counts[9] = 10

	counts := c.GetLastDaysCount(10)
	assert.Equal(t, counts, c.counts)

	counts = c.GetLastDaysCount(5)
	assert.Equal(t, counts, c.counts[5:])

	counts = c.GetLastDaysCount(1)
	assert.Equal(t, counts, c.counts[9:])

	counts = c.GetLastDaysCount(0)
	assert.Equal(t, counts, c.counts[9:])

	counts = c.GetLastDaysCount(-1)
	assert.Equal(t, counts, c.counts[9:])

	counts = c.GetLastDaysCount(11)
	assert.Equal(t, counts, c.counts)
}
