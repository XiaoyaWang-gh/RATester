package logger

import (
	"fmt"
	"testing"
)

func TestConfigDefault_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
