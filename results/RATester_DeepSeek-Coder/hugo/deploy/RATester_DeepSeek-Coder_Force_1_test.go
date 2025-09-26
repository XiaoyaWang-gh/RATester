package deploy

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestForce_1(t *testing.T) {
	type fields struct {
		NativePath string
		SlashPath  string
		UploadSize int64
		fs         afero.Fs
		matcher    *deployconfig.Matcher
		md5        []byte
		gzipped    bytes.Buffer
		mediaTypes media.Types
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "TestForce_True",
			fields: fields{
				matcher: &deployconfig.Matcher{Force: true},
			},
			want: true,
		},
		{
			name: "TestForce_False",
			fields: fields{
				matcher: &deployconfig.Matcher{Force: false},
			},
			want: false,
		},
		{
			name: "TestForce_Nil",
			fields: fields{
				matcher: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			lf := &localFile{
				NativePath: tt.fields.NativePath,
				SlashPath:  tt.fields.SlashPath,
				UploadSize: tt.fields.UploadSize,
				fs:         tt.fields.fs,
				matcher:    tt.fields.matcher,
				md5:        tt.fields.md5,
				gzipped:    tt.fields.gzipped,
				mediaTypes: tt.fields.mediaTypes,
			}
			if got := lf.Force(); got != tt.want {
				t.Errorf("localFile.Force() = %v, want %v", got, tt.want)
			}
		})
	}
}
