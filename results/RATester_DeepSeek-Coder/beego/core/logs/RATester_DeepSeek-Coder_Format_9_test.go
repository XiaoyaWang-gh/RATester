package logs

import (
	"fmt"
	"testing"
)

func TestFormat_9(t *testing.T) {
	tests := []struct {
		name string
		c    *connWriter
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
				t.Errorf("connWriter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
