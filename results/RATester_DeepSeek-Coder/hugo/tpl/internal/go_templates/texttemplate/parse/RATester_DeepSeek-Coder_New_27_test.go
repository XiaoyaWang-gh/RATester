package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_27(t *testing.T) {
	type args struct {
		name  string
		funcs []map[string]any
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "Test case 1",
			args: args{
				name:  "test",
				funcs: []map[string]any{{"func1": func() {}}},
			},
			want: &Tree{
				Name:  "test",
				funcs: []map[string]any{{"func1": func() {}}},
			},
		},
		{
			name: "Test case 2",
			args: args{
				name:  "test2",
				funcs: []map[string]any{{"func2": func() {}}},
			},
			want: &Tree{
				Name:  "test2",
				funcs: []map[string]any{{"func2": func() {}}},
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

			if got := New(tt.args.name, tt.args.funcs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
