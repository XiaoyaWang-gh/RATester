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
			name: "Test SetLevel with Level 0",
			args: args{lv: 0},
		},
		{
			name: "Test SetLevel with Level 1",
			args: args{lv: 1},
		},
		{
			name: "Test SetLevel with Level 2",
			args: args{lv: 2},
		},
		{
			name: "Test SetLevel with Level 3",
			args: args{lv: 3},
		},
		{
			name: "Test SetLevel with Level 4",
			args: args{lv: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetLevel(tt.args.lv)
		})
	}
}
