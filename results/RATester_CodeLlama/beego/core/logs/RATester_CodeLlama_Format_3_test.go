package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PatternLogFormatter{Pattern: "%F:%n|%w %t>> %m", WhenFormat: "2006-01-02"}
	lm := &LogMsg{Level: 1, Msg: "test", When: time.Now(), FilePath: "test.go", LineNumber: 100}
	if p.Format(lm) != "test.go:100|2019-01-02 15:04:05.000000000>> test" {
		t.Errorf("TestFormat failed")
	}
}
