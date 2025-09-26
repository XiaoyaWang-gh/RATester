package log

import (
	"fmt"
	"testing"
)

func TesttoString_1(t *testing.T) {
	tests := []struct {
		name string
		lv   Level
		want string
	}{
		{
			name: "Trace",
			lv:   LevelTrace,
			want: "Trace",
		},
		{
			name: "Debug",
			lv:   LevelDebug,
			want: "Debug",
		},
		{
			name: "Info",
			lv:   LevelInfo,
			want: "Info",
		},
		{
			name: "Warn",
			lv:   LevelWarn,
			want: "Warn",
		},
		{
			name: "Error",
			lv:   LevelError,
			want: "Error",
		},
		{
			name: "Invalid",
			lv:   Level(6),
			want: "[?6] ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.lv.toString(); got != tt.want {
				t.Errorf("Level.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
