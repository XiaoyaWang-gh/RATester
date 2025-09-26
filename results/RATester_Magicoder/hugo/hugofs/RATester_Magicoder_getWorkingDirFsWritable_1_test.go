package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestgetWorkingDirFsWritable_1(t *testing.T) {
	type args struct {
		base       afero.Fs
		workingDir string
	}
	tests := []struct {
		name string
		args args
		want afero.Fs
	}{
		{
			name: "Test Case 1",
			args: args{
				base:       afero.NewMemMapFs(),
				workingDir: "/test",
			},
			want: NewBasePathFs(afero.NewMemMapFs(), "/test"),
		},
		{
			name: "Test Case 2",
			args: args{
				base:       afero.NewMemMapFs(),
				workingDir: "",
			},
			want: afero.NewMemMapFs(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getWorkingDirFsWritable(tt.args.base, tt.args.workingDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWorkingDirFsWritable() = %v, want %v", got, tt.want)
			}
		})
	}
}
