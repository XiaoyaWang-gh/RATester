package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestStopDatadog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	datadogLoopCancelFunc = func() {}

	// when
	StopDatadog()

	// then
	assert.Nil(t, datadogLoopCancelFunc)
}
