package logs

import (
	"fmt"
	"testing"
)

func TestEnableFullFilePath_1(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestEnableFullFilePath_True",
			args: args{b: true},
		},
		{
			name: "TestEnableFullFilePath_False",
			args: args{b: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			EnableFullFilePath(tt.args.b)
		})
	}
}
