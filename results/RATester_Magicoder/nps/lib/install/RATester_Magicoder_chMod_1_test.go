package install

import (
	"fmt"
	"os"
	"testing"
)

func TestchMod_1(t *testing.T) {
	type args struct {
		name string
		mode os.FileMode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				name: "testfile",
				mode: 0644,
			},
		},
		{
			name: "Test case 2",
			args: args{
				name: "testdir",
				mode: os.ModeDir | 0755,
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

			chMod(tt.args.name, tt.args.mode)
		})
	}
}
