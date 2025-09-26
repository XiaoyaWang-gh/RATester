package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLogger_1(t *testing.T) {
	tests := []struct {
		name string
		args []int64
		want *BeeLogger
	}{
		{
			name: "Test with no arguments",
			args: []int64{},
			want: &BeeLogger{
				level:               LevelDebug,
				loggerFuncCallDepth: 3,
				msgChanLen:          defaultAsyncMsgLen,
			},
		},
		{
			name: "Test with one argument",
			args: []int64{100},
			want: &BeeLogger{
				level:               LevelDebug,
				loggerFuncCallDepth: 3,
				msgChanLen:          100,
			},
		},
		{
			name: "Test with zero argument",
			args: []int64{0},
			want: &BeeLogger{
				level:               LevelDebug,
				loggerFuncCallDepth: 3,
				msgChanLen:          defaultAsyncMsgLen,
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

			got := NewLogger(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
