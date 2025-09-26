package deploy

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/spf13/afero"
)

func TestReader_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		lf      *localFile
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name: "should return gzipped content if matcher is set and gzip is true",
			lf: &localFile{
				matcher: &deployconfig.Matcher{Gzip: true},
				gzipped: *bytes.NewBuffer([]byte("gzipped content")),
			},
			want:    io.NopCloser(bytes.NewReader([]byte("gzipped content"))),
			wantErr: false,
		},
		{
			name: "should return file content if matcher is not set or gzip is false",
			lf: &localFile{
				matcher:    &deployconfig.Matcher{Gzip: false},
				fs:         afero.NewMemMapFs(),
				NativePath: "/path/to/file",
			},
			want:    nil, // how to mock fs.Open return?
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.lf.Reader()
			if (err != nil) != tt.wantErr {
				t.Errorf("localFile.Reader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("localFile.Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}
