package logs

import (
	"fmt"
	"testing"
)

func TestMsgFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := []any{1, 2, 3}
	f := msgFunc(i...)
	if f() != "1 2 3" {
		t.Errorf("msgFunc() = %v, want %v", f(), "1 2 3")
	}
}
