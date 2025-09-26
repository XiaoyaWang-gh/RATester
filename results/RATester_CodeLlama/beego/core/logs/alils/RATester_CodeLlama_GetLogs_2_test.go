package alils

import (
	"fmt"
	"testing"
)

func TestGetLogs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add your test case here
}
