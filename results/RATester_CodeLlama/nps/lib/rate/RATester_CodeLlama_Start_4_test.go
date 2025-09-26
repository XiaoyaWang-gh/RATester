package rate

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestStart_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	rate := &Rate{
		bucketSize:        10,
		bucketSurplusSize: 10,
		stopChan:          make(chan bool),
	}
	// act
	rate.Start()
	// assert
	assert.NotNil(t, rate.stopChan)
}
