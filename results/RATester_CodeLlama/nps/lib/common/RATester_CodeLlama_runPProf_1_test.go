package common

import (
	"fmt"
	"testing"
)

func TestRunPProf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ipPort := "127.0.0.1:6060"
	runPProf(ipPort)
	// TODO: how to test the method?
}
