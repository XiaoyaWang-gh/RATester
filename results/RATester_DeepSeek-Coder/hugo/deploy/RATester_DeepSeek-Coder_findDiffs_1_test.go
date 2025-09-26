package deploy

import (
	"fmt"
	"reflect"
	"testing"

	"gocloud.dev/blob"
)

func TestFindDiffs_1(t *testing.T) {
	type args struct {
		localFiles  map[string]*localFile
		remoteFiles map[string]*blob.ListObject
		force       bool
	}
	tests := []struct {
		name  string
		args  args
		want  []*fileToUpload
		want1 []string
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

			d := &Deployer{}
			got, got1 := d.findDiffs(tt.args.localFiles, tt.args.remoteFiles, tt.args.force)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deployer.findDiffs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Deployer.findDiffs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
