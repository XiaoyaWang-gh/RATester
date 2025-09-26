package log

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInfof_1(t *testing.T) {
	t.Run("TestInfof", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		format := "test %v"
		v := "value"
		expected := fmt.Sprintf(format, v)
		buf := &bytes.Buffer{}
		logger.SetOutput(buf)
		Infof(format, v)
		got := buf.String()
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}
