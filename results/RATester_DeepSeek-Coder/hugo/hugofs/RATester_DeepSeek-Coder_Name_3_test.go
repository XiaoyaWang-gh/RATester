package hugofs

import (
	"fmt"
	"testing"
)

func TestName_3(t *testing.T) {
	type fields struct {
		noOpRegularFileOps *noOpRegularFileOps
		DirOnlyOps         DirOnlyOps
		fs                 *RootMappingFs
		name               string
		meta               *FileMeta
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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
			if got := f.Name(); got != tt.want {
				t.Errorf("rootMappingDir.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
