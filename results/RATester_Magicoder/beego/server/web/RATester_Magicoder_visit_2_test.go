package web

import (
	"fmt"
	"io/fs"
	"testing"
)

func Testvisit_2(t *testing.T) {
	type args struct {
		paths string
		f     fs.FileInfo
		err   error
	}
	tests := []struct {
		name    string
		tf      *templateFile
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

			if err := tt.tf.visit(tt.args.paths, tt.args.f, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("visit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
