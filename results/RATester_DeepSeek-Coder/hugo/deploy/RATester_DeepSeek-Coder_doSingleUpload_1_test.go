package deploy

import (
	"context"
	"fmt"
	"testing"

	"gocloud.dev/blob"
)

func TestDoSingleUpload_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		bucket *blob.Bucket
		upload *fileToUpload
	}
	tests := []struct {
		name    string
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

			d := &Deployer{}
			if err := d.doSingleUpload(tt.args.ctx, tt.args.bucket, tt.args.upload); (err != nil) != tt.wantErr {
				t.Errorf("Deployer.doSingleUpload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
