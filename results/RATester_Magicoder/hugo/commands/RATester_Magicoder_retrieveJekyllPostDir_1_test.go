package commands

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestretrieveJekyllPostDir_1(t *testing.T) {
	type args struct {
		fs  afero.Fs
		dir string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
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

			c := &importCommand{}
			got, got1 := c.retrieveJekyllPostDir(tt.args.fs, tt.args.dir)
			if got != tt.want {
				t.Errorf("retrieveJekyllPostDir() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("retrieveJekyllPostDir() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
