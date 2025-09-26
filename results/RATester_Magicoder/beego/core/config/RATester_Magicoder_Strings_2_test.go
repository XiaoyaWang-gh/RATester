package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_2(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				key: "section:key",
			},
			want:    []string{"val1", "val2"},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "section:key2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				key: "section:key3",
			},
			want:    []string{"val3"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &IniConfigContainer{}
			got, err := c.Strings(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
