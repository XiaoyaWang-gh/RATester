package logs

import (
	"fmt"
	"testing"
)

func TestWarning_3(t *testing.T) {
	bl := &BeeLogger{
		level: LevelWarn,
	}

	tests := []struct {
		name   string
		format string
		v      []interface{}
	}{
		{
			name:   "TestWarning",
			format: "This is a warning message",
			v:      []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			bl.Warning(tt.format, tt.v...)
		})
	}
}
