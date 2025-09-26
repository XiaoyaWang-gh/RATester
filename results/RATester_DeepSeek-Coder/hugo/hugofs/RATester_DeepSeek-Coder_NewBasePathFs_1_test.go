package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewBasePathFs_1(t *testing.T) {
	type args struct {
		source afero.Fs
		path   string
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

			if got := NewBasePathFs(tt.args.source, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBasePathFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
