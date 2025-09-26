package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
	"github.com/spf13/afero"
)

func TestHasPrefix_1(t *testing.T) {
	type fields struct {
		id            string
		Fs            afero.Fs
		rootMapToReal *radix.Tree
		realMapToRoot *radix.Tree
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fs.hasPrefix(tt.args.prefix); got != tt.want {
				t.Errorf("RootMappingFs.hasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
