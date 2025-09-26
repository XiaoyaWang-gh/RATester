package plugins

import (
	zipa "archive/zip"
	"fmt"
	"testing"
)

func TestUnzipFile_1(t *testing.T) {
	type args struct {
		f    *zipa.File
		dest string
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

			if err := unzipFile(tt.args.f, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("unzipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
