package deploy

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestWalkLocal_1(t *testing.T) {
	tests := []struct {
		name       string
		fs         afero.Fs
		matchers   []*deployconfig.Matcher
		include    glob.Glob
		exclude    glob.Glob
		mediaTypes media.Types
		mappath    func(string) string
		wantErr    bool
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

			d := &Deployer{
				localFs: tt.fs,
			}
			_, err := d.walkLocal(tt.fs, tt.matchers, tt.include, tt.exclude, tt.mediaTypes, tt.mappath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Deployer.walkLocal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
