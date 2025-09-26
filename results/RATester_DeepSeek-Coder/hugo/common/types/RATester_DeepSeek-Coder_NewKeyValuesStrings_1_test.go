package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewKeyValuesStrings_1(t *testing.T) {
	type args struct {
		key    string
		values []string
	}
	tests := []struct {
		name string
		args args
		want KeyValues
	}{
		{
			name: "Test case 1",
			args: args{
				key:    "testKey",
				values: []string{"value1", "value2"},
			},
			want: KeyValues{
				Key:    "testKey",
				Values: []any{"value1", "value2"},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewKeyValuesStrings(tt.args.key, tt.args.values...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyValuesStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
