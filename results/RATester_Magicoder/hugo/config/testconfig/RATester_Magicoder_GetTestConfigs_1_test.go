package testconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/config/allconfig"
	"github.com/spf13/afero"
)

func TestGetTestConfigs_1(t *testing.T) {
	type args struct {
		fs  afero.Fs
		cfg config.Provider
	}
	tests := []struct {
		name string
		args args
		want *allconfig.Configs
	}{
		{
			name: "Test with nil fs and cfg",
			args: args{
				fs:  nil,
				cfg: nil,
			},
			want: &allconfig.Configs{},
		},
		{
			name: "Test with fs and cfg",
			args: args{
				fs:  afero.NewMemMapFs(),
				cfg: config.New(),
			},
			want: &allconfig.Configs{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetTestConfigs(tt.args.fs, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
