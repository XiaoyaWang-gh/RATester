package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_4(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal []string
	}

	tests := []struct {
		name string
		c    *IniConfigContainer
		args args
		want []string
	}{
		{
			name: "TestDefaultStrings_0",
			c:    &IniConfigContainer{},
			args: args{
				key:        "key",
				defaultVal: []string{"default"},
			},
			want: []string{"default"},
		},
		{
			name: "TestDefaultStrings_1",
			c:    &IniConfigContainer{},
			args: args{
				key:        "key",
				defaultVal: []string{"default"},
			},
			want: []string{"default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.DefaultStrings(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
