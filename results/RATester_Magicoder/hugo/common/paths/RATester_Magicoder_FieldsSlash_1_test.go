package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFieldsSlash_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{s: "field/name"},
			want: []string{"field", "name"},
		},
		{
			name: "Test case 2",
			args: args{s: "field/name/with/multiple/slashes"},
			want: []string{"field", "name", "with", "multiple", "slashes"},
		},
		{
			name: "Test case 3",
			args: args{s: "field/name/with/trailing/slash/"},
			want: []string{"field", "name", "with", "trailing", "slash"},
		},
		{
			name: "Test case 4",
			args: args{s: "field/name/with/leading/slash"},
			want: []string{"field", "name", "with", "leading", "slash"},
		},
		{
			name: "Test case 5",
			args: args{s: "field/name/with/no/slashes"},
			want: []string{"field/name/with/no/slashes"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FieldsSlash(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldsSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
