package logs

import (
	"fmt"
	"testing"
)

func TestFormat_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SMTPWriter{}
	lm := &LogMsg{}
	got := s.Format(lm)
	want := lm.OldStyleFormat()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
