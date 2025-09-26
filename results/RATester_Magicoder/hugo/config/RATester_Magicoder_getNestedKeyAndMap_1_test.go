package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestgetNestedKeyAndMap_1(t *testing.T) {
	type args struct {
		key    string
		create bool
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 maps.Params
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

			c := &defaultConfigProvider{}
			got, got1 := c.getNestedKeyAndMap(tt.args.key, tt.args.create)
			if got != tt.want {
				t.Errorf("getNestedKeyAndMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getNestedKeyAndMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
