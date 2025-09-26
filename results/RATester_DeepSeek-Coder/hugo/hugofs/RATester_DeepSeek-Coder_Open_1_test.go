package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	radix "github.com/armon/go-radix"
	"github.com/spf13/afero"
)

func TestOpen_1(t *testing.T) {
	type fields struct {
		id            string
		Fs            afero.Fs
		rootMapToReal *radix.Tree
		realMapToRoot *radix.Tree
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    afero.File
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

			fs := &RootMappingFs{
				id:            tt.fields.id,
				Fs:            tt.fields.Fs,
				rootMapToReal: tt.fields.rootMapToReal,
				realMapToRoot: tt.fields.realMapToRoot,
			}
			got, err := fs.Open(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootMappingFs.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootMappingFs.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}
