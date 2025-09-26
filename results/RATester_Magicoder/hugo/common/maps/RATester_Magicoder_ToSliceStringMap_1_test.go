package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToSliceStringMap_1(t *testing.T) {
	tests := []struct {
		name    string
		in      any
		want    []map[string]any
		wantErr bool
	}{
		{
			name: "Test Case 1",
			in:   []map[string]any{{"key1": "value1", "key2": "value2"}},
			want: []map[string]any{{"key1": "value1", "key2": "value2"}},
		},
		{
			name: "Test Case 2",
			in:   Params{"key1": "value1", "key2": "value2"},
			want: []map[string]any{{"key1": "value1", "key2": "value2"}},
		},
		{
			name: "Test Case 3",
			in:   []any{map[string]any{"key1": "value1", "key2": "value2"}},
			want: []map[string]any{{"key1": "value1", "key2": "value2"}},
		},
		{
			name:    "Test Case 4",
			in:      "invalid input",
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

			got, err := ToSliceStringMap(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSliceStringMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSliceStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
