package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFileWriter_1(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		{
			name: "Test newFileWriter",
			want: &fileLogWriter{
				Daily:      true,
				MaxDays:    7,
				Hourly:     false,
				MaxHours:   168,
				Rotate:     true,
				RotatePerm: "0440",
				Level:      LevelTrace,
				Perm:       "0660",
				DirPerm:    "0770",
				MaxLines:   10000000,
				MaxFiles:   999,
				MaxSize:    1 << 28,
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

			if got := newFileWriter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFileWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
