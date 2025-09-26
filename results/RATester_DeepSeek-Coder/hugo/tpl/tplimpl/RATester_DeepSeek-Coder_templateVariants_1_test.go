package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTemplateVariants_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				name: "test",
			},
			want: []string{"test"},
		},
		{
			name: "Test case 2",
			args: args{
				name: "test.variant1",
			},
			want: []string{"variant1"},
		},
		{
			name: "Test case 3",
			args: args{
				name: "test.variant1.variant2",
			},
			want: []string{"variant1", "variant2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := templateVariants(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateVariants() = %v, want %v", got, tt.want)
			}
		})
	}
}
