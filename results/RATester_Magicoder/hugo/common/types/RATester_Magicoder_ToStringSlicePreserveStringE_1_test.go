package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringSlicePreserveStringE_1(t *testing.T) {
	tests := []struct {
		name    string
		v       any
		want    []string
		wantErr bool
	}{
		{
			name:    "nil",
			v:       nil,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "string",
			v:       "hello",
			want:    []string{"hello"},
			wantErr: false,
		},
		{
			name:    "slice of strings",
			v:       []string{"hello", "world"},
			want:    []string{"hello", "world"},
			wantErr: false,
		},
		{
			name:    "slice of ints",
			v:       []int{1, 2, 3},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
		{
			name:    "unsupported type",
			v:       123,
			want:    nil,
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

			got, err := ToStringSlicePreserveStringE(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringSlicePreserveStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlicePreserveStringE() = %v, want %v", got, tt.want)
			}
		})
	}
}
