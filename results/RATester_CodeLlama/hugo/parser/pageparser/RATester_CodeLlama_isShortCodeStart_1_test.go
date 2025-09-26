package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortCodeStart_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("{{<")
	l.pos = 0
	l.start = 0
	l.width = 0
	l.sectionHandlers = &sectionHandlers{}
	l.cfg = Config{}
	l.summaryDivider = []byte("<!--more-->")
	l.summaryDividerChecked = false
	l.lexerShortcodeState = lexerShortcodeState{}
	l.items = Items{}
	l.err = nil
	l.currLeftDelimItem = ItemType(0)
	l.currRightDelimItem = ItemType(0)
	l.isInline = false
	l.currShortcodeName = ""
	l.closingState = 0
	l.elementStepNum = 0
	l.paramElements = 0
	l.openShortcodes = map[string]bool{}
	l.run()
	if l.isShortCodeStart() != true {
		t.Errorf("isShortCodeStart() = %v, want %v", l.isShortCodeStart(), true)
	}
}
