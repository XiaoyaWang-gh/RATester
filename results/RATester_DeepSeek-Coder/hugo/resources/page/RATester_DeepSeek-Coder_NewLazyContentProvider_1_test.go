package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLazyContentProvider_1(t *testing.T) {
	type args struct {
		f func() (OutputFormatContentProvider, error)
	}
	tests := []struct {
		name string
		args args
		want *LazyContentProvider
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

			if got := NewLazyContentProvider(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLazyContentProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
