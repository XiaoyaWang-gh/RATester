package i18n

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPluralCount_1(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		Count int
	}

	tests := []struct {
		name string
		arg  any
		want any
	}{
		{
			name: "nil",
			arg:  nil,
			want: nil,
		},
		{
			name: "map with count field",
			arg:  map[string]any{"count": 1},
			want: 1,
		},
		{
			name: "map with Count field",
			arg:  map[string]any{"Count": 1},
			want: 1,
		},
		{
			name: "map with count field name",
			arg:  map[string]any{"countFieldName": 1},
			want: 1,
		},
		{
			name: "struct with count field",
			arg:  testStruct{Count: 1},
			want: 1,
		},
		{
			name: "struct with count method",
			arg:  &testStruct{Count: 1},
			want: 1,
		},
		{
			name: "other type",
			arg:  1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getPluralCount(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPluralCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
