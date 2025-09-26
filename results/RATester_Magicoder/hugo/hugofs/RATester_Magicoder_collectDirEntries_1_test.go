package hugofs

import (
	"fmt"
	iofs "io/fs"
	"reflect"
	"testing"
)

func TestcollectDirEntries_1(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		rfs     *RootMappingFs
		args    args
		want    []iofs.DirEntry
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

			got, err := tt.rfs.collectDirEntries(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootMappingFs.collectDirEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootMappingFs.collectDirEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}
