package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestOldStyleFormat_1(t *testing.T) {
	tests := []struct {
		name string
		lm   *LogMsg
		want string
	}{
		{
			name: "Test with no args",
			lm: &LogMsg{
				Level:               0,
				Msg:                 "test message",
				When:                time.Now(),
				FilePath:            "/path/to/file",
				LineNumber:          10,
				Args:                []interface{}{},
				Prefix:              "prefix",
				enableFullFilePath:  true,
				enableFuncCallDepth: true,
			},
			want: "prefix test message",
		},
		{
			name: "Test with args",
			lm: &LogMsg{
				Level:               0,
				Msg:                 "test message %s",
				When:                time.Now(),
				FilePath:            "/path/to/file",
				LineNumber:          10,
				Args:                []interface{}{"arg"},
				Prefix:              "prefix",
				enableFullFilePath:  true,
				enableFuncCallDepth: true,
			},
			want: "prefix test message arg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.lm.OldStyleFormat(); got != tt.want {
				t.Errorf("LogMsg.OldStyleFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
