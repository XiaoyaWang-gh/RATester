package compress

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconfigDefault_16(t *testing.T) {
	tests := []struct {
		name string
		args []Config
		want Config
	}{
		{
			name: "Test Case 1",
			args: []Config{
				{
					Level: LevelDefault,
				},
			},
			want: Config{
				Level: LevelDefault,
			},
		},
		{
			name: "Test Case 2",
			args: []Config{
				{
					Level: LevelDisabled,
				},
			},
			want: Config{
				Level: LevelDefault,
			},
		},
		{
			name: "Test Case 3",
			args: []Config{
				{
					Level: LevelBestCompression,
				},
			},
			want: Config{
				Level: LevelBestCompression,
			},
		},
		{
			name: "Test Case 4",
			args: []Config{
				{
					Level: LevelBestSpeed,
				},
			},
			want: Config{
				Level: LevelBestSpeed,
			},
		},
		{
			name: "Test Case 5",
			args: []Config{
				{
					Level: 3,
				},
			},
			want: Config{
				Level: LevelDefault,
			},
		},
		{
			name: "Test Case 6",
			args: []Config{
				{
					Level: -2,
				},
			},
			want: Config{
				Level: LevelDefault,
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

			if got := configDefault(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
