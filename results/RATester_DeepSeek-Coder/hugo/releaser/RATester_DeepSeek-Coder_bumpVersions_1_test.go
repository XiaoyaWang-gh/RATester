package releaser

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugo"
)

func TestBumpVersions_1(t *testing.T) {
	type args struct {
		ver hugo.Version
	}
	tests := []struct {
		name    string
		r       *ReleaseHandler
		args    args
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

			if err := tt.r.bumpVersions(tt.args.ver); (err != nil) != tt.wantErr {
				t.Errorf("ReleaseHandler.bumpVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
