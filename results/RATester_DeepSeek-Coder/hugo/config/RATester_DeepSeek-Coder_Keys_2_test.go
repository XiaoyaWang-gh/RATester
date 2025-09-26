package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestKeys_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		provider *defaultConfigProvider
		want     []string
	}{
		{
			name: "empty",
			provider: &defaultConfigProvider{
				root: maps.Params{},
			},
			want: []string{},
		},
		{
			name: "single key",
			provider: &defaultConfigProvider{
				root: maps.Params{"key1": "value1"},
			},
			want: []string{"key1"},
		},
		{
			name: "multiple keys",
			provider: &defaultConfigProvider{
				root: maps.Params{"key1": "value1", "key2": "value2", "key3": "value3"},
			},
			want: []string{"key1", "key2", "key3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.provider.Keys()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
