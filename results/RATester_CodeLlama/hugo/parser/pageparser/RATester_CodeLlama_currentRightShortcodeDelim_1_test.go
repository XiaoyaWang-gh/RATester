package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCurrentRightShortcodeDelim_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.currRightDelimItem = tRightDelimScWithMarkup
	if !reflect.DeepEqual(l.currentRightShortcodeDelim(), rightDelimScWithMarkup) {
		t.Errorf("currentRightShortcodeDelim() = %v, want %v", l.currentRightShortcodeDelim(), rightDelimScWithMarkup)
	}
}
