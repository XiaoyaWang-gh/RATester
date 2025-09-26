package xlog

import (
	"fmt"
	"testing"
)

func TestInfof_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &Logger{}
	logger.prefixString = "test"

	format := "test format"
	v := []interface{}{"test value"}

	logger.Infof(format, v...)

	// assert something here
}
