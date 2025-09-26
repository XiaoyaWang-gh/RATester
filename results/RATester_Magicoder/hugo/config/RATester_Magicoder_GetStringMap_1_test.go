package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetStringMap_1(t *testing.T) {
	type fields struct {
		root maps.Params
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]any
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

			c := &defaultConfigProvider{
				root: tt.fields.root,
			}
			if got := c.GetStringMap(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultConfigProvider.GetStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
