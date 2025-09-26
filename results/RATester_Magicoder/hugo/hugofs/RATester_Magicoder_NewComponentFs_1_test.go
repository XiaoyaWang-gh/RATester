package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/spf13/afero"
)

func TestNewComponentFs_1(t *testing.T) {
	type args struct {
		opts ComponentFsOptions
	}
	tests := []struct {
		name string
		args args
		want *componentFs
	}{
		{
			name: "Test case 1",
			args: args{
				opts: ComponentFsOptions{
					Fs:                     afero.NewMemMapFs(),
					Component:              "content",
					DefaultContentLanguage: "en",
					PathParser:             &paths.PathParser{},
				},
			},
			want: &componentFs{
				Fs: afero.NewMemMapFs(),
				opts: ComponentFsOptions{
					Fs:                     afero.NewMemMapFs(),
					Component:              "content",
					DefaultContentLanguage: "en",
					PathParser:             &paths.PathParser{},
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				opts: ComponentFsOptions{
					Fs:                     afero.NewMemMapFs(),
					Component:              "layouts",
					DefaultContentLanguage: "fr",
					PathParser:             &paths.PathParser{},
				},
			},
			want: &componentFs{
				Fs: afero.NewMemMapFs(),
				opts: ComponentFsOptions{
					Fs:                     afero.NewMemMapFs(),
					Component:              "layouts",
					DefaultContentLanguage: "fr",
					PathParser:             &paths.PathParser{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewComponentFs(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComponentFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
