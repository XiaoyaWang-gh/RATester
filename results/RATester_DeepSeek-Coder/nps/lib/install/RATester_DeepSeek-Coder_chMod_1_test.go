package install

import (
	"fmt"
	"os"
	"testing"
)

func TestChMod_1(t *testing.T) {
	type args struct {
		name string
		mode os.FileMode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with a valid file and mode",
			args: args{
				name: "testfile",
				mode: 0755,
			},
		},
		{
			name: "Test with a valid file and mode",
			args: args{
				name: "testfile",
				mode: 0644,
			},
		},
		{
			name: "Test with a non-existing file",
			args: args{
				name: "nonexistentfile",
				mode: 0777,
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
			fileInfo, err := os.Stat(tt.args.name)
			if err != nil {
				t.Errorf("Error in getting file info: %v", err)
			}
			if fileInfo.Mode() != tt.args.mode {
				t.Errorf("Expected mode %v, got %v", tt.args.mode, fileInfo.Mode())
			}
		})
	}
}
