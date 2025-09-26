package server

import (
	"fmt"
	"testing"
)

func TestStopMetricsClients_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	// when
	stopMetricsClients()
	// then
	// assert
}
