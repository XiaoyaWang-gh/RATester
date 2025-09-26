package loggers

import (
	"fmt"
	"testing"
)

func TestidfInfoStatement_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	what := "test"
	id := "test_id"
	format := "test_format"
	expected := fmt.Sprintf("\nYou can suppress this %s by adding the following to your site configuration:\nignoreLogs = ['%s']", what, id)
	result := l.idfInfoStatement(what, id, format)
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
