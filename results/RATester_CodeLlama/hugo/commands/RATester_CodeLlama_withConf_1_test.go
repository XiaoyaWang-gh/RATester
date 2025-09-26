package commands

import (
	"fmt"
	"testing"
)

func TestWithConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	c := &hugoBuilder{}
	fn := func(conf *commonConfig) {}

	// When
	c.withConf(fn)

	// Then
	// ...
}
