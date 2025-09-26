package hugofs

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

func TestUnwrapFilesystem_1(t *testing.T) {
	type fields struct {
		Fs afero.Fs
		re *regexp.Regexp
	}
	tests := []struct {
		name   string
		fields fields
		want   afero.Fs
	}{
		{
			name: "TestUnwrapFilesystem",
			fields: fields{
				Fs: afero.NewMemMapFs(),
				re: regexp.MustCompile(".*"),
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

			fs := &stacktracerFs{
				Fs: tt.fields.Fs,
				re: tt.fields.re,
			}
			if got := fs.UnwrapFilesystem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stacktracerFs.UnwrapFilesystem() = %v, want %v", got, tt.want)
			}
		})
	}
}
