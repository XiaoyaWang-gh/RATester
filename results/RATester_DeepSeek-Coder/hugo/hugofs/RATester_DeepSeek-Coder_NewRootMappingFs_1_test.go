package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewRootMappingFs_1(t *testing.T) {
	type args struct {
		fs  afero.Fs
		rms []RootMapping
	}
	tests := []struct {
		name    string
		args    args
		want    *RootMappingFs
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

			got, err := NewRootMappingFs(tt.args.fs, tt.args.rms...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRootMappingFs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRootMappingFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
