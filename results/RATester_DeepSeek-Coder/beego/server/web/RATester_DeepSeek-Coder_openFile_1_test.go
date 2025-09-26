package web

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestOpenFile_1(t *testing.T) {
	type args struct {
		filePath       string
		fi             os.FileInfo
		acceptEncoding string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   string
		want2   *serveContentHolder
		want3   *serveContentReader
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

			got, got1, got2, got3, err := openFile(tt.args.filePath, tt.args.fi, tt.args.acceptEncoding)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("openFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("openFile() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("openFile() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("openFile() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
