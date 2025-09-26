package logs

import (
	"fmt"
	"testing"
)

func TestNewLogger_1(t *testing.T) {
	tests := []struct {
		name string
		want *BeeLogger
	}{
		{
			name: "Test NewLogger",
			want: &BeeLogger{
				level:               LevelDebug,
				loggerFuncCallDepth: 3,
				msgChanLen:          0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewLogger()
			if got.level != tt.want.level {
				t.Errorf("NewLogger() = %v, want %v", got.level, tt.want.level)
			}
			if got.loggerFuncCallDepth != tt.want.loggerFuncCallDepth {
				t.Errorf("NewLogger() = %v, want %v", got.loggerFuncCallDepth, tt.want.loggerFuncCallDepth)
			}
			if got.msgChanLen != tt.want.msgChanLen {
				t.Errorf("NewLogger() = %v, want %v", got.msgChanLen, tt.want.msgChanLen)
			}
		})
	}
}
