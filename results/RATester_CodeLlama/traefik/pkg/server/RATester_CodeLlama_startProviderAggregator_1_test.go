package server

import (
	"fmt"
	"testing"
)

func TestStartProviderAggregator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &ConfigurationWatcher{}

	// when
	c.startProviderAggregator()

	// then
	// TODO
}
