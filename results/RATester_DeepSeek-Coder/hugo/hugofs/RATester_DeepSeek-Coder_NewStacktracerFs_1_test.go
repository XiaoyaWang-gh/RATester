package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewStacktracerFs_1(t *testing.T) {
	type args struct {
		fs      afero.Fs
		pattern string
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

			if got := NewStacktracerFs(tt.args.fs, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStacktracerFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
