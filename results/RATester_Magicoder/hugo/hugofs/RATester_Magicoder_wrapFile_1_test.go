package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestwrapFile_1(t *testing.T) {
	type args struct {
		f afero.File
	}
	tests := []struct {
		name string
		args args
		want afero.File
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

			fs := &hashingFs{}
			if got := fs.wrapFile(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashingFs.wrapFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
