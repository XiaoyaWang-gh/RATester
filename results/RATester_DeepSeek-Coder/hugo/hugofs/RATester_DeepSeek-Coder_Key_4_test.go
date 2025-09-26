package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
	"github.com/spf13/afero"
)

func TestKey_4(t *testing.T) {
	type fields struct {
		id            string
		Fs            afero.Fs
		rootMapToReal *radix.Tree
		realMapToRoot *radix.Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				id:            "testID",
				Fs:            nil,
				rootMapToReal: nil,
				realMapToRoot: nil,
			},
			want: "testID",
		},
		{
			name: "Test Case 2",
			fields: fields{
				id:            "anotherID",
				Fs:            nil,
				rootMapToReal: nil,
				realMapToRoot: nil,
			},
			want: "anotherID",
		},
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
			if got := fs.Key(); got != tt.want {
				t.Errorf("RootMappingFs.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
