package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testlength_1(t *testing.T) {
	tests := []struct {
		name    string
		item    reflect.Value
		want    int
		wantErr bool
	}{
		{
			name:    "Test length of array",
			item:    reflect.ValueOf([]int{1, 2, 3, 4, 5}),
			want:    5,
			wantErr: false,
		},
		{
			name:    "Test length of string",
			item:    reflect.ValueOf("hello"),
			want:    5,
			wantErr: false,
		},
		{
			name:    "Test length of nil pointer",
			item:    reflect.ValueOf(nil),
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test length of unsupported type",
			item:    reflect.ValueOf(123),
			want:    0,
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

			got, err := length(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("length() = %v, want %v", got, tt.want)
			}
		})
	}
}
