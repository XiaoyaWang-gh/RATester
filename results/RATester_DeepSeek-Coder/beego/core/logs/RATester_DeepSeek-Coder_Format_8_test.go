package logs

import (
	"fmt"
	"testing"
)

func TestFormat_8(t *testing.T) {
	tests := []struct {
		name string
		c    *consoleWriter
		lm   *LogMsg
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.Format(tt.lm); got != tt.want {
				t.Errorf("consoleWriter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
