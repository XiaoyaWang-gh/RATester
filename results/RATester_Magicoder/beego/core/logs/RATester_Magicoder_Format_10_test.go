package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	jlw := &JLWriter{
		AuthorName: "test",
		Title:      "test",
	}
	lm := &LogMsg{
		Msg:  "test",
		When: time.Now(),
	}
	expected := fmt.Sprintf("%s %s", lm.When.Format("2006-01-02 15:04:05"), lm.Msg)
	result := jlw.Format(lm)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
