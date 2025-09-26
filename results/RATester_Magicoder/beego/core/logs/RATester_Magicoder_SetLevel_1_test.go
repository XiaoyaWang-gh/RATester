package logs

import (
	"fmt"
	"testing"
)

func TestSetLevel_1(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSetLevel_0",
			args: args{l: 0},
		},
		{
			name: "TestSetLevel_1",
			args: args{l: 1},
		},
		{
			name: "TestSetLevel_2",
			args: args{l: 2},
		},
		{
			name: "TestSetLevel_3",
			args: args{l: 3},
		},
		{
			name: "TestSetLevel_4",
			args: args{l: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetLevel(tt.args.l)
		})
	}
}
