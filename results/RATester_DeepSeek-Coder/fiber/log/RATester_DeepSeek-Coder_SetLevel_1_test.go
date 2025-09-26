package log

import (
	"fmt"
	"testing"
)

func TestSetLevel_1(t *testing.T) {
	type args struct {
		lv Level
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test SetLevel with LevelDebug",
			args: args{
				lv: LevelDebug,
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

			SetLevel(tt.args.lv)
			// Add assertions to check if the level was set correctly
		})
	}
}
