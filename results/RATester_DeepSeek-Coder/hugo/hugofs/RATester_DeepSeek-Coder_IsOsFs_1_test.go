package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsOsFs_1(t *testing.T) {
	type args struct {
		fs afero.Fs
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with OsFs",
			args: args{
				fs: &afero.OsFs{},
			},
			want: true,
		},
		{
			name: "Test with MemMapFs",
			args: args{
				fs: &afero.MemMapFs{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsOsFs(tt.args.fs); got != tt.want {
				t.Errorf("IsOsFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
