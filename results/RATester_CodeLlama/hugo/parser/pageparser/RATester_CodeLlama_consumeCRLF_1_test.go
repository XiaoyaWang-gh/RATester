package pageparser

import (
	"fmt"
	"testing"
)

func TestConsumeCRLF_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("abc\r\n")
	l.pos = 0
	l.start = 0
	l.width = 0
	l.state = l.stateStart
	l.stateStart = l.stateStart
	l.sectionHandlers = &sectionHandlers{}
	l.cfg = Config{}
	l.summaryDivider = []byte("")
	l.summaryDividerChecked = false
	l.lexerShortcodeState = lexerShortcodeState{}
	l.items = Items{}
	l.err = nil
	l.currLeftDelimItem = 0
	l.currRightDelimItem = 0
	l.isInline = false
	l.currShortcodeName = ""
	l.closingState = 0
	l.elementStepNum = 0
	l.paramElements = 0
	l.openShortcodes = map[string]bool{}
	l.consumeCRLF()
	if l.pos != 4 {
		t.Errorf("l.pos = %d, want 4", l.pos)
	}
	if l.width != 2 {
		t.Errorf("l.width = %d, want 2", l.width)
	}
}
