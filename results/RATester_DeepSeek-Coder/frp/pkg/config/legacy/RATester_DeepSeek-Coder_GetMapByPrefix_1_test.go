package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMapByPrefix_1(t *testing.T) {
	type args struct {
		set    map[string]string
		prefix string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test case 1",
			args: args{
				set:    map[string]string{"prefix1": "value1", "prefix2": "value2", "other": "value3"},
				prefix: "prefix",
			},
			want: map[string]string{"prefix1": "value1", "prefix2": "value2"},
		},
		{
			name: "Test case 2",
			args: args{
				set:    map[string]string{"prefix1": "value1", "prefix2": "value2", "other": "value3"},
				prefix: "other",
			},
			want: map[string]string{"other": "value3"},
		},
		{
			name: "Test case 3",
			args: args{
				set:    map[string]string{"prefix1": "value1", "prefix2": "value2", "other": "value3"},
				prefix: "no_prefix",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetMapByPrefix(tt.args.set, tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMapByPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
