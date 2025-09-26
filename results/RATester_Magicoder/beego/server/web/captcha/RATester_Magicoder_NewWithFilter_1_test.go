package captcha

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWithFilter_1(t *testing.T) {
	type args struct {
		urlPrefix string
		store     Storage
	}
	tests := []struct {
		name string
		args args
		want *Captcha
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

			if got := NewWithFilter(tt.args.urlPrefix, tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
