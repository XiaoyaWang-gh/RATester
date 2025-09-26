package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestGetJekyllDirInfo_1(t *testing.T) {
	type args struct {
		fs         afero.Fs
		jekyllRoot string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]bool
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
			got, got1 := c.getJekyllDirInfo(tt.args.fs, tt.args.jekyllRoot)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getJekyllDirInfo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getJekyllDirInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
