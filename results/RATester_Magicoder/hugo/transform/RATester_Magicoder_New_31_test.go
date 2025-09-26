package transform

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_31(t *testing.T) {
	type args struct {
		trs []Transformer
	}
	tests := []struct {
		name string
		args args
		want Chain
	}{
		{
			name: "Test case 1",
			args: args{
				trs: []Transformer{
					func(ft FromTo) error {
						// Implementation of Transformer
						return nil
					},
					func(ft FromTo) error {
						// Implementation of Transformer
						return nil
					},
				},
			},
			want: []Transformer{
				func(ft FromTo) error {
					// Implementation of Transformer
					return nil
				},
				func(ft FromTo) error {
					// Implementation of Transformer
					return nil
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

			if got := New(tt.args.trs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
