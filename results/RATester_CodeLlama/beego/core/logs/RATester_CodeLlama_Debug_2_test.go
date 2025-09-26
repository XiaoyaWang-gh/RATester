package logs

import (
	"fmt"
	"testing"
)

func TestDebug_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := "test"
	v := []interface{}{"test"}
	Debug(f, v...)
}
