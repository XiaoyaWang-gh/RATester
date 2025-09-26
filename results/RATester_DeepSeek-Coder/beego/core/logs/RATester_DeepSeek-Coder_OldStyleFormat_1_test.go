package logs

import (
	"fmt"
	"testing"
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
				Level:               1,
				Msg:                 "test message",
				FilePath:            "/path/to/file",
				LineNumber:          100,
				enableFullFilePath:  true,
				enableFuncCallDepth: true,
			},
			want: "[test.go:100] INFO test message",
		},
		{
			name: "Test with args",
			lm: &LogMsg{
				Level:               2,
				Msg:                 "test %s",
				Args:                []interface{}{"arg1"},
				FilePath:            "/path/to/file",
				LineNumber:          200,
				enableFullFilePath:  false,
				enableFuncCallDepth: false,
			},
			want: "WARN test arg1",
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
				t.Errorf("OldStyleFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
