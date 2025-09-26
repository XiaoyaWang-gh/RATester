package logs

import (
	"fmt"
	"testing"
)

func TestDelLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := NewLogger()
	bl.SetLogger("console", "{\"level\":1}")
	bl.DelLogger("console")
	if len(bl.outputs) != 0 {
		t.Errorf("DelLogger failed")
	}
}
