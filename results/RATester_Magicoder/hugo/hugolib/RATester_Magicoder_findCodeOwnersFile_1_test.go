package hugolib

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestfindCodeOwnersFile_1(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
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

			got, err := findCodeOwnersFile(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("findCodeOwnersFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCodeOwnersFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
