package echo

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestFileFS_1(t *testing.T) {
	type args struct {
		file       string
		filesystem fs.FS
	}
	tests := []struct {
		name    string
		c       *context
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

			if err := tt.c.FileFS(tt.args.file, tt.args.filesystem); (err != nil) != tt.wantErr {
				t.Errorf("context.FileFS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
