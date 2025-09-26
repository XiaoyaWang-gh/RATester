package glob

import (
	"fmt"
	"testing"
)

func TestDoMatch_1(t *testing.T) {
	type args struct {
		filename string
		isDir    bool
	}
	tests := []struct {
		name string
		f    *FilenameFilter
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.doMatch(tt.args.filename, tt.args.isDir); got != tt.want {
				t.Errorf("doMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
