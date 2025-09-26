package web

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func Testwalk_1(t *testing.T) {
	type args struct {
		fs     http.FileSystem
		path   string
		info   os.FileInfo
		walkFn filepath.WalkFunc
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

			if err := walk(tt.args.fs, tt.args.path, tt.args.info, tt.args.walkFn); (err != nil) != tt.wantErr {
				t.Errorf("walk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
