package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetNestedKeyAndMap_1(t *testing.T) {
	type args struct {
		key    string
		create bool
	}
	tests := []struct {
		name  string
		c     *defaultConfigProvider
		args  args
		want  string
		want1 maps.Params
	}{
		{
			name: "test1",
			c: &defaultConfigProvider{
				root: maps.Params{
					"a": maps.Params{
						"b": "c",
					},
				},
			},
			args: args{
				key:    "a.b",
				create: true,
			},
			want: "b",
			want1: maps.Params{
				"c": "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := tt.c.getNestedKeyAndMap(tt.args.key, tt.args.create)
			if got != tt.want {
				t.Errorf("getNestedKeyAndMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getNestedKeyAndMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
