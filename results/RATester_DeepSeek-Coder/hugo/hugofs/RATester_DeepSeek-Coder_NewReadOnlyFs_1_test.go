package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewReadOnlyFs_1(t *testing.T) {
	type args struct {
		source afero.Fs
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

			if got := NewReadOnlyFs(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReadOnlyFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
