package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestnewRootMappingFsFromFromTo_1(t *testing.T) {
	type args struct {
		baseDir string
		fs      afero.Fs
		fromTo  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *RootMappingFs
		wantErr bool
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

			got, err := newRootMappingFsFromFromTo(tt.args.baseDir, tt.args.fs, tt.args.fromTo...)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRootMappingFsFromFromTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRootMappingFsFromFromTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
