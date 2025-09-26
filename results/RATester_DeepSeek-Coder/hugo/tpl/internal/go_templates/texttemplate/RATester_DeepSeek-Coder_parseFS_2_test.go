package template

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestParseFS_2(t *testing.T) {
	type args struct {
		fsys     fs.FS
		patterns []string
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

			_, err := parseFS(nil, tt.args.fsys, tt.args.patterns)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
