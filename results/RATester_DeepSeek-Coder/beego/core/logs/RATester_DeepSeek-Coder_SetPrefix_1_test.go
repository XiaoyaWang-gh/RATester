package logs

import (
	"fmt"
	"testing"
)

func TestSetPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &BeeLogger{prefix: "oldPrefix"}
	newPrefix := "newPrefix"
	logger.SetPrefix(newPrefix)
	if logger.prefix != newPrefix {
		t.Errorf("Expected prefix to be %s, but got %s", newPrefix, logger.prefix)
	}
}
