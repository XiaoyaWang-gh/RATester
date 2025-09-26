package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestWalkFilesystems_1(t *testing.T) {
	type args struct {
		fs afero.Fs
		fn WalkFn
	}

	tests := []struct {
		name string
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

			if got := WalkFilesystems(tt.args.fs, tt.args.fn); got != tt.want {
				t.Errorf("WalkFilesystems() = %v, want %v", got, tt.want)
			}
		})
	}
}
