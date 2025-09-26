package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDictionary_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		values  []any
		want    map[string]any
		wantErr bool
	}{
		{
			name:   "simple dictionary",
			values: []any{"key1", "value1", "key2", "value2"},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name:   "nested dictionary",
			values: []any{"key1", "value1", "key2", "value2", "key3", "value3"},
			want: map[string]any{
				"key1": "value1",
				"key2": map[string]any{
					"key3": "value3",
				},
			},
			wantErr: false,
		},
		{
			name:    "invalid dictionary call",
			values:  []any{"key1", "value1", "key2"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid dictionary key",
			values:  []any{123, "value1", "key2", "value2"},
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

			got, err := ns.Dictionary(tt.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dictionary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dictionary() = %v, want %v", got, tt.want)
			}
		})
	}
}
