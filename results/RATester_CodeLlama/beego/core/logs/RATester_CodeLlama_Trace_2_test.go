package logs

import (
	"fmt"
	"testing"
)

func TestTrace_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := "hello %s"
	v := []interface{}{"world"}
	beeLogger.Trace(formatPattern(f, v...), v...)
}
