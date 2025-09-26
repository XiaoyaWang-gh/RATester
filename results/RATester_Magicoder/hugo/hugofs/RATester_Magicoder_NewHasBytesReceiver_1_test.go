package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewHasBytesReceiver_1(t *testing.T) {
	type args struct {
		delegate         afero.Fs
		shouldCheck      func(name string) bool
		hasBytesCallback func(name string, match []byte)
		patterns         [][]byte
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

			if got := NewHasBytesReceiver(tt.args.delegate, tt.args.shouldCheck, tt.args.hasBytesCallback, tt.args.patterns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHasBytesReceiver() = %v, want %v", got, tt.want)
			}
		})
	}
}
