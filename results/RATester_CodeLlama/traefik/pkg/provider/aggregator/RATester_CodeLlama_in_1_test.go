package aggregator

import (
	"fmt"
	"testing"
)

func TestIn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ch := &RingChannel{}
	ch.in()
}
