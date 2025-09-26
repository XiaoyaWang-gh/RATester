package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestNewFrom_1(t *testing.T) {
	type args struct {
		fs   afero.Fs
		conf config.BaseConfig
	}
	tests := []struct {
		name string
		args args
		want *Fs
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

			if got := NewFrom(tt.args.fs, tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
