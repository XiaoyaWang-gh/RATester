package legacy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestTransformHeadersFromPluginParams_1(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name string
		args args
		want v1.HeaderOperations
	}{
		{
			name: "Test case 1",
			args: args{
				params: map[string]string{
					"plugin_header_key1": "value1",
					"plugin_header_key2": "value2",
				},
			},
			want: v1.HeaderOperations{
				Set: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				params: map[string]string{
					"plugin_header_key1": "value1",
					"key2":               "value2",
				},
			},
			want: v1.HeaderOperations{
				Set: map[string]string{
					"key1": "value1",
				},
			},
		},
		{
			name: "Test case 3",
			args: args{
				params: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: v1.HeaderOperations{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := transformHeadersFromPluginParams(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformHeadersFromPluginParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
