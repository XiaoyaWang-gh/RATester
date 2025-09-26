package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoFloat_1(t *testing.T) {
	tests := []struct {
		name    string
		v       reflect.Value
		want    float64
		wantErr bool
	}{
		{
			name:    "float32",
			v:       reflect.ValueOf(float32(1.23)),
			want:    1.23,
			wantErr: false,
		},
		{
			name:    "float64",
			v:       reflect.ValueOf(float64(4.56)),
			want:    4.56,
			wantErr: false,
		},
		{
			name:    "int",
			v:       reflect.ValueOf(int(789)),
			want:    789,
			wantErr: false,
		},
		{
			name:    "invalid",
			v:       reflect.ValueOf("invalid"),
			want:    -1,
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

			got, err := toFloat(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
