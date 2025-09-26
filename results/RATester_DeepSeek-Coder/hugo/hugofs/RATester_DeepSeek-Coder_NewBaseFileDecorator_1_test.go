package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewBaseFileDecorator_1(t *testing.T) {
	type args struct {
		fs        afero.Fs
		callbacks []func(fi FileMetaInfo)
	}
	tests := []struct {
		name string
		args args
		want afero.Fs
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

			if got := NewBaseFileDecorator(tt.args.fs, tt.args.callbacks...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseFileDecorator() = %v, want %v", got, tt.want)
			}
		})
	}
}
