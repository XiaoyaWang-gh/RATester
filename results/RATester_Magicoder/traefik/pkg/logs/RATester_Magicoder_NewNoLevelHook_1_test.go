package logs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewNoLevelHook_1(t *testing.T) {
	type args struct {
		minLevel zerolog.Level
		level    zerolog.Level
	}
	tests := []struct {
		name string
		args args
		want *NoLevelHook
	}{
		{
			name: "Test case 1",
			args: args{
				minLevel: zerolog.DebugLevel,
				level:    zerolog.InfoLevel,
			},
			want: &NoLevelHook{
				minLevel: zerolog.DebugLevel,
				level:    zerolog.InfoLevel,
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewNoLevelHook(tt.args.minLevel, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoLevelHook() = %v, want %v", got, tt.want)
			}
		})
	}
}
