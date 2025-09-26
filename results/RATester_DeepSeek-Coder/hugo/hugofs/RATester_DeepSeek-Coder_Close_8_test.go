package hugofs

import (
	"fmt"
	"testing"
)

func TestClose_8(t *testing.T) {
	type fields struct {
		noOpRegularFileOps *noOpRegularFileOps
		DirOnlyOps         DirOnlyOps
		fs                 *RootMappingFs
		name               string
		meta               *FileMeta
	}
	tests := []struct {
		name    string
		fields  fields
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

			f := &rootMappingDir{
				noOpRegularFileOps: tt.fields.noOpRegularFileOps,
				DirOnlyOps:         tt.fields.DirOnlyOps,
				fs:                 tt.fields.fs,
				name:               tt.fields.name,
				meta:               tt.fields.meta,
			}
			if err := f.Close(); (err != nil) != tt.wantErr {
				t.Errorf("rootMappingDir.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
