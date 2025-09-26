package resources

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestopenPublishFileForWriting_1(t *testing.T) {
	type args struct {
		relTargetPath string
	}
	tests := []struct {
		name    string
		r       *genericResource
		args    args
		want    io.WriteCloser
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

			got, err := tt.r.openPublishFileForWriting(tt.args.relTargetPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("genericResource.openPublishFileForWriting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.openPublishFileForWriting() = %v, want %v", got, tt.want)
			}
		})
	}
}
