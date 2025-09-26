package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestParams_3(t *testing.T) {
	type fields struct {
		params map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		want   maps.Params
	}{
		{
			name: "Test case 1",
			fields: fields{
				params: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: maps.Params{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Test case 2",
			fields: fields{
				params: map[string]any{
					"key3": "value3",
					"key4": "value4",
				},
			},
			want: maps.Params{
				"key3": "value3",
				"key4": "value4",
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

			l := &genericResource{
				params: tt.fields.params,
			}
			if got := l.Params(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.Params() = %v, want %v", got, tt.want)
			}
		})
	}
}
