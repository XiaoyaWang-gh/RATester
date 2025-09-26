package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_2(t *testing.T) {
	t.Parallel()

	type args struct {
		key string
	}

	tests := []struct {
		name    string
		c       *IniConfigContainer
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "TestStrings_0",
			c:    &IniConfigContainer{},
			args: args{
				key: "key0",
			},
			want:    []string{"val0", "val1"},
			wantErr: false,
		},
		{
			name: "TestStrings_1",
			c:    &IniConfigContainer{},
			args: args{
				key: "key1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "TestStrings_2",
			c:    &IniConfigContainer{},
			args: args{
				key: "key2",
			},
			want:    []string{"val0"},
			wantErr: false,
		},
		{
			name: "TestStrings_3",
			c:    &IniConfigContainer{},
			args: args{
				key: "key3",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.c.Strings(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() got = %v, want %v", got, tt.want)
			}
		})
	}
}
