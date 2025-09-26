package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestWrapFilesystem_1(t *testing.T) {
	type args struct {
		container afero.Fs
		content   afero.Fs
	}
	tests := []struct {
		name string
		args args
		want afero.Fs
	}{
		{
			name: "Test case 1",
			args: args{
				container: afero.NewMemMapFs(),
				content:   afero.NewMemMapFs(),
			},
			want: filesystemsWrapper{
				Fs:      afero.NewMemMapFs(),
				content: afero.NewMemMapFs(),
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WrapFilesystem(tt.args.container, tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapFilesystem() = %v, want %v", got, tt.want)
			}
		})
	}
}
