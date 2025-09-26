package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnumberToFloat_1(t *testing.T) {
	tests := []struct {
		name    string
		v       reflect.Value
		want    float64
		wantErr bool
	}{
		{
			name:    "Test with float64",
			v:       reflect.ValueOf(float64(1.23)),
			want:    1.23,
			wantErr: false,
		},
		{
			name:    "Test with int",
			v:       reflect.ValueOf(int(123)),
			want:    123,
			wantErr: false,
		},
		{
			name:    "Test with uint",
			v:       reflect.ValueOf(uint(123)),
			want:    123,
			wantErr: false,
		},
		{
			name:    "Test with invalid kind",
			v:       reflect.ValueOf("invalid"),
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

			got, err := numberToFloat(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("numberToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("numberToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
