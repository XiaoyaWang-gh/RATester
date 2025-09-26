package logs

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PatternLogFormatter{Pattern: "%F:%n|%w %t>> %m", WhenFormat: "2006-01-02"}
	lm := &LogMsg{Level: 1, Msg: "test", When: time.Now(), FilePath: "test.go", LineNumber: 100, Args: []interface{}{"test"}}
	res := p.ToString(lm)
	assert.Equal(t, "test.go:100|2022-01-01 12:00:00 AM>> test", res)
}
