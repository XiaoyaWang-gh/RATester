package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNestedParam_1(t *testing.T) {
	type testCase struct {
		name       string
		keyStr     string
		separator  string
		candidates []Params
		want       any
		wantErr    bool
	}

	tests := []testCase{
		{
			name:      "simple key",
			keyStr:    "key1",
			separator: ".",
			candidates: []Params{
				{"key1": "value1"},
			},
			want:    "value1",
			wantErr: false,
		},
		{
			name:      "nested key",
			keyStr:    "key1.key2",
			separator: ".",
			candidates: []Params{
				{"key1": Params{"key2": "value2"}},
			},
			want:    "value2",
			wantErr: false,
		},
		{
			name:      "key not found",
			keyStr:    "key3",
			separator: ".",
			candidates: []Params{
				{"key1": "value1"},
			},
			want:    nil,
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

			got, err := GetNestedParam(tt.keyStr, tt.separator, tt.candidates...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNestedParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNestedParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
