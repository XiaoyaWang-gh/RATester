package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_4(t *testing.T) {
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
				key: "test_key",
			},
			want:    []string{"test_value1", "test_value2"},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "non_existing_key",
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

			got, err := Strings(tt.args.key)
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
