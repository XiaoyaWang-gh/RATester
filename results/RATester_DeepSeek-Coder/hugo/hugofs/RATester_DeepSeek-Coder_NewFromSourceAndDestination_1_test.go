package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestNewFromSourceAndDestination_1(t *testing.T) {
	type args struct {
		source      afero.Fs
		destination afero.Fs
		cfg         config.Provider
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

			if got := NewFromSourceAndDestination(tt.args.source, tt.args.destination, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromSourceAndDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
