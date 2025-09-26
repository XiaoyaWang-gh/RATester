package logs

import (
	"fmt"
	"testing"
)

func TestNewConsole_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
