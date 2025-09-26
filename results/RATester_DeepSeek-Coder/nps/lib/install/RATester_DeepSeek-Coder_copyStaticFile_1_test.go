package install

import (
	"fmt"
	"testing"
)

func TestCopyStaticFile_1(t *testing.T) {
	type args struct {
		srcPath string
		bin     string
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := copyStaticFile(tt.args.srcPath, tt.args.bin); got != tt.want {
				t.Errorf("copyStaticFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
