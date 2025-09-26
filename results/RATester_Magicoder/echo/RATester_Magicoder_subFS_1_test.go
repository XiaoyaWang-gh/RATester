package echo

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestSubFS_1(t *testing.T) {
	type args struct {
		currentFs fs.FS
		root      string
	}
	tests := []struct {
		name    string
		args    args
		want    fs.FS
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

			got, err := subFS(tt.args.currentFs, tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("subFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
