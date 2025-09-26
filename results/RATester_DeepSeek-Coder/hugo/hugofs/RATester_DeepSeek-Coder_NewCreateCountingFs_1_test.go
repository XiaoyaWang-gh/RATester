package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewCreateCountingFs_1(t *testing.T) {
	type args struct {
		fs afero.Fs
	}
	tests := []struct {
		name string
		args args
		want afero.Fs
	}{
		{
			name: "TestNewCreateCountingFs",
			args: args{fs: afero.NewMemMapFs()},
			want: &createCountingFs{Fs: afero.NewMemMapFs(), fileCount: make(map[string]int)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewCreateCountingFs(tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateCountingFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
