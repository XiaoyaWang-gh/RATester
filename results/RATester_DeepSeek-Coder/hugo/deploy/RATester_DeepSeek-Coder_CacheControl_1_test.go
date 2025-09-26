package deploy

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestCacheControl_1(t *testing.T) {
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
		want   string
	}{
		{
			name: "Test with nil matcher",
			fields: fields{
				matcher: nil,
			},
			want: "",
		},
		{
			name: "Test with non-nil matcher",
			fields: fields{
				matcher: &deployconfig.Matcher{
					CacheControl: "public, max-age=3600",
				},
			},
			want: "public, max-age=3600",
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
			if got := lf.CacheControl(); got != tt.want {
				t.Errorf("localFile.CacheControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
