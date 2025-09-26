package alils

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestFormat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &aliLSWriter{}
	lm := &logs.LogMsg{
		Msg: "test message",
	}
	expected := lm.OldStyleFormat()
	actual := w.Format(lm)
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
