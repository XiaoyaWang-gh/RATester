package pagemeta

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestsetParamIfNotSet_1(t *testing.T) {
	type args struct {
		key   string
		value any
		d     *FrontMatterDescriptor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				key:   "testKey",
				value: "testValue",
				d:     &FrontMatterDescriptor{PageConfig: &PageConfig{Params: maps.Params{}}},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				key:   "testKey",
				value: "testValue",
				d:     &FrontMatterDescriptor{PageConfig: &PageConfig{Params: maps.Params{"testKey": "existingValue"}}},
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

			setParamIfNotSet(tt.args.key, tt.args.value, tt.args.d)
			if _, found := tt.args.d.PageConfig.Params[tt.args.key]; !found {
				t.Errorf("setParamIfNotSet() failed to set key %s", tt.args.key)
			}
			if tt.args.d.PageConfig.Params[tt.args.key] != tt.args.value {
				t.Errorf("setParamIfNotSet() failed to set correct value for key %s", tt.args.key)
			}
		})
	}
}
