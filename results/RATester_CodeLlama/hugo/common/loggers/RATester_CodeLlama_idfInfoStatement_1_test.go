package loggers

import (
	"fmt"
	"testing"
)

func TestIdfInfoStatement_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	what := "error"
	id := "1234567890"
	format := "You can suppress this %s by adding the following to your site configuration:\nignoreLogs = ['%s']"
	want := "\nYou can suppress this error by adding the following to your site configuration:\nignoreLogs = ['1234567890']"
	got := l.idfInfoStatement(what, id, format)
	if got != want {
		t.Errorf("idfInfoStatement() = %q, want %q", got, want)
	}
}
