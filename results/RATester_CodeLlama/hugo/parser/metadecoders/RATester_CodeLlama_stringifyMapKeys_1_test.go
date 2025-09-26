package metadecoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringifyMapKeys_1(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "test_case_1",
			args: args{
				in: map[any]any{
					"a": map[any]any{
						"b": map[any]any{
							"c": "d",
						},
					},
				},
			},
			want: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": "d",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got, _ := stringifyMapKeys(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringifyMapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
