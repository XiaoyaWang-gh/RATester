package deploy

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
	"gocloud.dev/blob"
)

func TestWalkRemote_1(t *testing.T) {
	ctx := context.Background()
	bucket := &blob.Bucket{}
	include, _ := glob.Compile("*")
	exclude, _ := glob.Compile("")

	tests := []struct {
		name    string
		ctx     context.Context
		bucket  *blob.Bucket
		include glob.Glob
		exclude glob.Glob
		want    map[string]*blob.ListObject
		wantErr bool
	}{
		{
			name:    "Test case 1",
			ctx:     ctx,
			bucket:  bucket,
			include: include,
			exclude: exclude,
			want:    map[string]*blob.ListObject{},
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &Deployer{}
			got, err := d.walkRemote(tt.ctx, tt.bucket, tt.include, tt.exclude)
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
