package collections

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_55(t *testing.T) {
	type args struct {
		deps *deps.Deps
	}
	tests := []struct {
		name string
		args args
		want *Namespace
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

			if got := New(tt.args.deps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
