package deploy

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
	"gocloud.dev/blob"
)

func TestwalkRemote_1(t *testing.T) {
	ctx := context.Background()
	bucket := &blob.Bucket{}
	include := glob.MustCompile("*")
	exclude := glob.MustCompile("*")

	tests := []struct {
		name    string
		bucket  *blob.Bucket
		include glob.Glob
		exclude glob.Glob
		want    map[string]*blob.ListObject
		wantErr bool
	}{
		{
			name:    "test case 1",
			bucket:  bucket,
			include: include,
			exclude: exclude,
			want:    map[string]*blob.ListObject{},
			wantErr: false,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &Deployer{}
			got, err := d.walkRemote(ctx, tt.bucket, tt.include, tt.exclude)
			if (err != nil) != tt.wantErr {
				t.Errorf("walkRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walkRemote() = %v, want %v", got, tt.want)
			}
		})
	}
}
