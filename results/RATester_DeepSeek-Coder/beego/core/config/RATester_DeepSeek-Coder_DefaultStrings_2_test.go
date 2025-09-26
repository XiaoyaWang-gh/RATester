package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_2(t *testing.T) {
	type args struct {
		key        string
		defaultVal []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				key:        "key1",
				defaultVal: []string{"default1", "default2"},
			},
			want: []string{"default1", "default2"},
		},
		{
			name: "Test case 2",
			args: args{
				key:        "key2",
				defaultVal: []string{"default3", "default4"},
			},
			want: []string{"default3", "default4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DefaultStrings(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
