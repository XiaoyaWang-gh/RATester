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

	bl := &BeeLogger{}
	prefix := "test_prefix"
	bl.SetPrefix(prefix)
	if bl.prefix != prefix {
		t.Errorf("Expected prefix to be %s, but got %s", prefix, bl.prefix)
	}
}
