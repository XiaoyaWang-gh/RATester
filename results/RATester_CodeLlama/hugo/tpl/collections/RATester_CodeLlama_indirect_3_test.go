package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirect_3(t *testing.T) {
	tests := []struct {
		name  string
		v     reflect.Value
		want  reflect.Value
		want1 bool
	}{
		{
			name:  "test case 1",
			v:     reflect.ValueOf(nil),
			want:  reflect.ValueOf(nil),
			want1: true,
		},
		{
			name:  "test case 2",
			v:     reflect.ValueOf(1),
			want:  reflect.ValueOf(1),
			want1: false,
		},
		{
			name:  "test case 3",
			v:     reflect.ValueOf(1),
			want:  reflect.ValueOf(1),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := indirect(tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirect() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("indirect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
