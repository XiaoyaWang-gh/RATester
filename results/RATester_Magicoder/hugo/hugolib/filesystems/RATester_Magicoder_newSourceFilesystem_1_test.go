package filesystems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestnewSourceFilesystem_1(t *testing.T) {
	type args struct {
		name string
		fs   afero.Fs
	}
	tests := []struct {
		name string
		args args
		want *SourceFilesystem
	}{
		{
			name: "Test Case 1",
			args: args{
				name: "test",
				fs:   afero.NewMemMapFs(),
			},
			want: &SourceFilesystem{
				Name:     "test",
				Fs:       afero.NewMemMapFs(),
				SourceFs: afero.NewMemMapFs(),
			},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &sourceFilesystemsBuilder{
				sourceFs: afero.NewMemMapFs(),
			}
			if got := b.newSourceFilesystem(tt.args.name, tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSourceFilesystem() = %v, want %v", got, tt.want)
			}
		})
	}
}
