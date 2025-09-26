package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLookupEqualFold_1(t *testing.T) {
	type args struct {
		m   map[string]any
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  any
		want1 string
		want2 bool
	}{
		{
			name: "Test case 1",
			args: args{
				m:   map[string]any{"Hello": "World", "Go": "Gopher"},
				key: "hello",
			},
			want:  "World",
			want1: "Hello",
			want2: true,
		},
		{
			name: "Test case 2",
			args: args{
				m:   map[string]any{"Hello": "World", "Go": "Gopher"},
				key: "Go",
			},
			want:  "Gopher",
			want1: "Go",
			want2: true,
		},
		{
			name: "Test case 3",
			args: args{
				m:   map[string]any{"Hello": "World", "Go": "Gopher"},
				key: "NonExistent",
			},
			want:  nil,
			want1: "",
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1, got2 := LookupEqualFold(tt.args.m, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) || got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("LookupEqualFold() = %v, %v, %v, want %v, %v, %v", got, got1, got2, tt.want, tt.want1, tt.want2)
			}
		})
	}
}
