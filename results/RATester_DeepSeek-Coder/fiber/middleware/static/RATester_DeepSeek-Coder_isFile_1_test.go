package static

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestIsFile_1(t *testing.T) {
	type args struct {
		root       string
		filesystem fs.FS
	}
	tests := []struct {
		name    string
		args    args
		want    bool
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

			got, err := isFile(tt.args.root, tt.args.filesystem)
			if (err != nil) != tt.wantErr {
				t.Errorf("isFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
